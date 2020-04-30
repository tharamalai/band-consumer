package meicdp

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	"github.com/tharamalai/meichain/x/meicdp/types"
)

// NewHandler creates the msg handler of this module, as required by Cosmos-SDK standard.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case channeltypes.MsgPacket:
			var responseData oracle.OracleResponsePacketData
			if err := types.ModuleCdc.UnmarshalJSON(msg.GetData(), &responseData); err == nil {
				return handleOracleResponsePacketData(ctx, keeper, responseData)
			}

			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"cannot unmarshal oracle packet data",
			)

		case MsgLockCollateral:
			return handleMsgLockCollateral(ctx, keeper, msg)

		case MsgUnlockCollateral:
			return handleOracleRequestPacketData(ctx, keeper, msg, msg.Sender)

		case MsgBorrowDebt:
			return handleOracleRequestPacketData(ctx, keeper, msg, msg.Sender)

		case MsgReturnDebt:
			return handleMsgReturnDebt(ctx, keeper, msg)

		case MsgSetSourceChannel:
			// TODO: Check permission
			return handleSetSourceChannel(ctx, keeper, msg)

		default:
			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %T",
				ModuleName,
				msg,
			)
		}
	}
}

func handleMsgLockCollateral(ctx sdk.Context, keeper Keeper, msg MsgLockCollateral) (*sdk.Result, error) {

	channelID, err := keeper.GetChannel(ctx, CosmosHubChain, "transfer")
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidChannel,
			fmt.Sprintf("channel of %s chain not found", CosmosHubChain),
		)
	}

	denom := fmt.Sprintf("transfer/%s/%s", channelID, types.AtomUnit)

	cdp := keeper.GetCDP(ctx, msg.Sender)

	lockCoin := sdk.NewCoin(denom, sdk.NewInt(int64(msg.Amount)))
	lockAmountCoins := sdk.NewCoins(lockCoin)

	collateralCoin := sdk.NewCoin(denom, sdk.NewInt(int64(cdp.CollateralAmount)))

	//  Accumulate collateral on CDP
	collateralCoin = collateralCoin.Add(lockCoin)
	if collateralCoin.IsNegative() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid lock amount. collateral must more than or equals 0.",
		)
	}

	cdp.CollateralAmount = collateralCoin.Amount.Uint64()

	// Transfer collateral to the module account. Transaction fails if sender's balance is insufficient.
	err = keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, lockAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"can't transfer %s coins from sender to CDP",
			denom,
		)
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgReturnDebt(ctx sdk.Context, keeper Keeper, msg MsgReturnDebt) (*sdk.Result, error) {

	cdp := keeper.GetCDP(ctx, msg.Sender)

	// Subtract debt on CDP
	returnCoin := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(msg.Amount)))
	returnAmountCoins := sdk.NewCoins(returnCoin)

	debtCoin := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(cdp.DebtAmount)))

	// New debt
	debtCoin = debtCoin.Sub(returnCoin)
	if debtCoin.IsNegative() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid return amount. debt must more than or equals 0.",
		)
	}
	cdp.DebtAmount = debtCoin.Amount.Uint64()

	// TODO: Pay fee

	// Transfer Mei from user to CDP. Transaction fails if sender's balance is insufficient.
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, sdkerrors.ErrInsufficientFunds
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// CDP burn returned coins
	err = keeper.SupplyKeeper.BurnCoins(ctx, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, types.ErrBurnCoin
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleSetSourceChannel(ctx sdk.Context, keeper Keeper, msg types.MsgSetSourceChannel) (*sdk.Result, error) {
	keeper.SetChannel(ctx, msg.ChainName, msg.SourcePort, msg.SourceChannel)
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func requestOracle(ctx sdk.Context, keeper Keeper, dataReq DataRequest) error {

	channelID, err := keeper.GetChannel(ctx, dataReq.ChainID, dataReq.Port)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidChannel,
			"channel %s not found",
			dataReq.Port,
		)
	}

	sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, dataReq.Port, channelID)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port meichain",
			dataReq.Port,
		)
	}

	destinationPort := sourceChannelEnd.Counterparty.PortID
	destinationChannel := sourceChannelEnd.Counterparty.ChannelID
	sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
		ctx, dataReq.Port, channelID,
	)

	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			dataReq.Port,
		)
	}

	packet := oracle.NewOracleRequestPacketData(
		dataReq.ClientID, dataReq.OracleScriptID, dataReq.Calldata,
		dataReq.AskCount, dataReq.MinCount,
	)

	return keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
		sequence, dataReq.Port, channelID, destinationPort, destinationChannel,
		1000000000, // Arbitrarily high timeout for now
	))

}

func handleOracleRequestPacketData(ctx sdk.Context, keeper Keeper, msg sdk.Msg, sender sdk.AccAddress) (*sdk.Result, error) {
	msgID := keeper.GetNextMsgCount(ctx)

	// Setup oracle request
	bandChainID := BandChainID
	port := ModuleName
	oracleScriptID := oracle.OracleScriptID(OracleScriptID)
	clientID := fmt.Sprintf("Msg:%d", msgID)
	calldata := encodeRequestParams(AtomSymbol, AtomMultiplier)
	askCount := int64(1)
	minCount := int64(1)

	channelID, err := keeper.GetChannel(ctx, bandChainID, port)

	dataRequest := types.NewDataRequest(
		oracleScriptID,
		channelID,
		bandChainID,
		port,
		clientID,
		calldata,
		askCount,
		minCount,
		sender,
	)

	// Set message to the store for waiting the oracle response packet.
	keeper.SetMsg(ctx, msgID, msg)

	err = requestOracle(ctx, keeper, dataRequest)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrRequestOracleData,
			"error while request oracle data: %v",
			err,
		)
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleOracleResponsePacketData(ctx sdk.Context, keeper Keeper, packet oracle.OracleResponsePacketData) (*sdk.Result, error) {

	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return nil, sdkerrors.Wrapf(
			types.ErrUnknownClientID,
			"unknown client id %s",
			packet.ClientID,
		)
	}

	msgID, err := strconv.ParseUint(clientID[1], 10, 64)
	if err != nil {
		return nil, err
	}

	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return nil, err
	}

	decoder := types.NewDecoder(rawResult)
	collateralPrice, err := decoder.DecodeU64()
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"cannot decode orable result data",
		)
	}

	msg, err := keeper.GetMsg(ctx, msgID)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrMsgNotFound,
			"cannot get stored message",
		)
	}

	switch msg := msg.(type) {
	case MsgUnlockCollateral:
		return handleMsgUnlockCollateral(ctx, keeper, msg, collateralPrice)

	case MsgBorrowDebt:
		return handleMsgBorrowDebt(ctx, keeper, msg, collateralPrice)

	default:
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidMsgType,
			fmt.Sprintf("invalid message type: %T", msg),
		)
	}

}

// handleMsgUnlockCollateral handles the unlock collateral message after receives oracle packet
func handleMsgUnlockCollateral(ctx sdk.Context, keeper Keeper, msg MsgUnlockCollateral, collateralPrice uint64) (*sdk.Result, error) {

	cosmosHubChannelID, err := keeper.GetChannel(ctx, CosmosHubChain, "transfer")
	denom := fmt.Sprintf("transfer/%s/%s", cosmosHubChannelID, types.AtomUnit)

	cdp := keeper.GetCDP(ctx, msg.Sender)

	unlockCoin := sdk.NewCoin(denom, sdk.NewInt(int64(msg.Amount)))
	unlockAmountCoins := sdk.NewCoins(unlockCoin)

	collateralCoin := sdk.NewCoin(denom, sdk.NewInt(int64(cdp.CollateralAmount)))

	// Subtract collateral on CDP
	collateralCoin = collateralCoin.Sub(unlockCoin)
	if collateralCoin.IsNegative() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid unlock amount. collateral must more than or equals 0.",
		)
	}

	cdp.CollateralAmount = collateralCoin.Amount.Uint64()

	// Calculate new collateral ratio. If collateral is lower than 150 percent then returns error.
	debtAmount := new(big.Int).SetUint64(cdp.DebtAmount)
	minimumDebtAmount := new(big.Int).SetUint64(0)
	if debtAmount.Cmp(minimumDebtAmount) > 0 {

		collateralRatioFloat := calculateCollateralRatioOfCDP(cdp, collateralPrice, AtomMultiplier)
		minimunRatio := new(big.Float).SetFloat64(MinimumCollateralRatio)
		collateralRatio, _ := collateralRatioFloat.Float64()
		if collateralRatioFloat.Cmp(minimunRatio) == -1 {
			return nil, sdkerrors.Wrapf(
				types.ErrTooLowCollateralRatio,
				fmt.Sprintf("collateral ratio is too low. (%.2f)", collateralRatio),
			)
		}
	}

	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, msg.Sender, unlockAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "can't transfer coins from CDP module to sender")
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

// handleMsgBorrowDebt handles the borrow debt message after receives oracle packet
func handleMsgBorrowDebt(ctx sdk.Context, keeper Keeper, msg types.MsgBorrowDebt, collateralPrice uint64) (*sdk.Result, error) {
	cdp := keeper.GetCDP(ctx, msg.Sender)

	borrowCoin := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(msg.Amount)))
	borrowAmountCoins := sdk.NewCoins(borrowCoin)

	// Accumurate debt on CDP
	debtCoin := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(cdp.DebtAmount)))
	debtCoin = debtCoin.Add(borrowCoin)

	if debtCoin.IsNegative() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid borrow amount. debt must more than or equals 0.",
		)
	}

	cdp.DebtAmount = debtCoin.Amount.Uint64()

	// Calculate new collateral ratio. If collateral is lower than 150 percent then returns error.
	debtAmount := new(big.Int).SetUint64(cdp.DebtAmount)
	minimumDebtAmount := new(big.Int).SetUint64(0)
	if debtAmount.Cmp(minimumDebtAmount) > 0 {

		collateralRatioFloat := calculateCollateralRatioOfCDP(cdp, collateralPrice, AtomMultiplier)
		minimunRatio := new(big.Float).SetFloat64(MinimumCollateralRatio)
		collateralRatio, _ := collateralRatioFloat.Float64()
		if collateralRatioFloat.Cmp(minimunRatio) == -1 {
			return nil, sdkerrors.Wrapf(
				types.ErrTooLowCollateralRatio,
				fmt.Sprintf("collateral ratio is too low. (%.2f)", collateralRatio),
			)
		}
	}

	// CDP mint Mei coins
	err := keeper.SupplyKeeper.MintCoins(ctx, ModuleName, borrowAmountCoins)
	if err != nil {
		return nil, types.ErrMintCoin
	}

	// CDP send coins from module to sender
	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, msg.Sender, borrowAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"can't transfer coins from CDP module to sender",
		)
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
