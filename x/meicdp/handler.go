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
				err := handleOracleRespondPacketData(ctx, keeper, responseData)
				if err != nil {
					return nil, sdkerrors.Wrapf(
						types.ErrResponseOracleData,
						"error while handle response oracle data: %v",
						err,
					)
				}

				return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
			}
			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"cannot unmarshal oracle packet data",
			)

		case MsgLockCollateral:
			return handleMsgLockCollateral(ctx, keeper, msg)

		case MsgUnlockCollateral:
			err := handleOracleRequestPacketData(ctx, keeper, msg, msg.Sender)
			if err != nil {
				return nil, sdkerrors.Wrapf(
					types.ErrResponseOracleData,
					"error while handle request oracle data: %v",
					err,
				)
			}
			return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

		case MsgBorrowDebt:
			err := handleOracleRequestPacketData(ctx, keeper, msg, msg.Sender)
			if err != nil {
				return nil, sdkerrors.Wrapf(
					types.ErrResponseOracleData,
					"error while handle request oracle data: %v",
					err,
				)
			}
			return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

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

	lockAmount := sdk.NewCoin(denom, sdk.NewInt(int64(msg.Amount)))
	lockAmountCoins := sdk.NewCoins(lockAmount)

	//  Accumulate collateral on CDP
	lockAmountInt := new(big.Int).SetUint64(msg.Amount)
	collateralAmountInt := new(big.Int).SetUint64(cdp.CollateralAmount)
	collateralAmountInt.Add(collateralAmountInt, lockAmountInt)
	if !collateralAmountInt.IsUint64() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid lock amount. collateral must more than or equals 0.",
		)
	}

	cdp.CollateralAmount = collateralAmountInt.Uint64()

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

	return &sdk.Result{}, nil
}

func handleMsgReturnDebt(ctx sdk.Context, keeper Keeper, msg MsgReturnDebt) (*sdk.Result, error) {

	cdp := keeper.GetCDP(ctx, msg.Sender)

	// Subtract debt on CDP
	returnAmount := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(msg.Amount)))
	returnAmountCoins := sdk.NewCoins(returnAmount)

	// New debt
	returnAmountInt := new(big.Int).SetUint64(msg.Amount)
	debtAmountInt := new(big.Int).SetUint64(cdp.DebtAmount)
	debtAmountInt.Sub(debtAmountInt, returnAmountInt)
	if !debtAmountInt.IsUint64() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid return amount. debt must more than or equals 0.",
		)
	}
	cdp.DebtAmount = debtAmountInt.Uint64()

	// TODO: Pay fee

	// Transfer Mei from user to CDP. Transaction fails if sender's balance is insufficient.
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"can't transfer coins from sender to CDP",
		)
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// CDP burn returned coins
	err = keeper.SupplyKeeper.BurnCoins(ctx, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrBurnCoin,
			"burn coin fail",
		)
	}

	return &sdk.Result{}, nil
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

func handleOracleRequestPacketData(ctx sdk.Context, keeper Keeper, msg sdk.Msg, sender sdk.AccAddress) error {
	msgID := keeper.GetMsgCount(ctx)

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
		return sdkerrors.Wrapf(
			types.ErrRequestOracleData,
			"error while request oracle data: %v",
			err,
		)
	}

	return nil
}

func handleOracleRespondPacketData(ctx sdk.Context, keeper Keeper, packet oracle.OracleResponsePacketData) error {

	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return sdkerrors.Wrapf(
			types.ErrUnknownClientID,
			"unknown client id %s",
			packet.ClientID,
		)
	}

	msgID, err := strconv.ParseUint(clientID[1], 10, 64)
	if err != nil {
		return err
	}

	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return err
	}

	decoder := types.NewDecoder(rawResult)
	collateralPrice, err := decoder.DecodeU64()
	if err != nil {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"cannot decode orable result data",
		)
	}

	msg, err := keeper.GetMsg(ctx, msgID)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMsgNotFound,
			"cannot get stored message",
		)
	}

	switch msg := msg.(type) {
	case MsgUnlockCollateral:
		err := handleMsgUnlockCollateral(ctx, keeper, msg, collateralPrice)
		if err != nil {
			return err
		}

	case MsgBorrowDebt:
		err := handleMsgBorrowDebt(ctx, keeper, msg, collateralPrice)
		if err != nil {
			return err
		}
	}

	default:
		return sdkerrors.Wrapf(
			types.ErrInvalidMsgType,
			fmt.Sprintf("invalid message type: %T", msg),
		)

	return nil

}

// handleMsgUnlockCollateral handles the unlock collateral message after receives oracle packet
func handleMsgUnlockCollateral(ctx sdk.Context, keeper Keeper, msg MsgUnlockCollateral, collateralPrice uint64) error {

	cosmosHubChannelID, err := keeper.GetChannel(ctx, CosmosHubChain, "transfer")
	denom := fmt.Sprintf("transfer/%s/%s", cosmosHubChannelID, types.AtomUnit)

	cdp := keeper.GetCDP(ctx, msg.Sender)

	unlockAmount := sdk.NewCoin(denom, sdk.NewInt(int64(msg.Amount)))
	unlockAmountCoins := sdk.NewCoins(unlockAmount)

	// Subtract collateral on CDP
	unlockAmountInt := new(big.Int).SetUint64(msg.Amount)
	collateralAmount := new(big.Int).SetUint64(cdp.CollateralAmount)
	collateralAmount.Sub(collateralAmount, unlockAmountInt)

	minimumCollateralAmount := new(big.Int).SetUint64(0)
	if collateralAmount.Cmp(minimumCollateralAmount) == -1 {
		return sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid unlock amount. collateral must more than or equals 0.",
		)
	}

	cdp.CollateralAmount = collateralAmount.Uint64()

	// Calculate new collateral ratio. If collateral is lower than 150 percent then returns error.
	debtAmount := new(big.Int).SetUint64(cdp.DebtAmount)
	minimumDebtAmount := new(big.Int).SetUint64(0)
	if debtAmount.Cmp(minimumDebtAmount) > 0 {

		collateralRatioFloat := calculateCollateralRatioOfCDP(cdp, collateralPrice, AtomMultiplier)
		minimunRatio := new(big.Float).SetFloat64(MinimumCollateralRatio)
		collateralRatio, _ := collateralRatioFloat.Float64()
		if collateralRatioFloat.Cmp(minimunRatio) == -1 {
			return sdkerrors.Wrapf(
				types.ErrTooLowCollateralRatio,
				fmt.Sprintf("collateral rate is too low. (%f%)", collateralRatio),
			)
		}
	}

	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, msg.Sender, unlockAmountCoins)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "can't transfer coins from CDP module to sender")
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	return nil
}

// handleMsgBorrowDebt handles the borrow debt message after receives oracle packet
func handleMsgBorrowDebt(ctx sdk.Context, keeper Keeper, msg types.MsgBorrowDebt, collateralPrice uint64) error {
	cdp := keeper.GetCDP(ctx, msg.Sender)

	borrowAmount := sdk.NewCoin(types.MeiUnit, sdk.NewInt(int64(msg.Amount)))
	borrowAmountCoins := sdk.NewCoins(borrowAmount)

	// Accumurate debt on CDP
	borrowAmountInt := new(big.Int).SetUint64(msg.Amount)
	debtAmountUint64 := new(big.Int).SetUint64(cdp.DebtAmount)
	debtAmountUint64.Add(debtAmountUint64, borrowAmountInt)
	minimumDebtAmount := new(big.Int).SetUint64(0)
	if debtAmountUint64.Cmp(minimumDebtAmount) == -1 {
		return sdkerrors.Wrapf(
			types.ErrInvalidBasicMsg,
			"invalid borrow amount. debt must more than or equals 0.",
		)
	}

	cdp.DebtAmount = debtAmountUint64.Uint64()

	// Calculate new collateral ratio. If collateral is lower than 150 percent then returns error.
	debtAmount := new(big.Int).SetUint64(cdp.DebtAmount)
	if debtAmount.Cmp(minimumDebtAmount) > 0 {

		collateralRatioFloat := calculateCollateralRatioOfCDP(cdp, collateralPrice, AtomMultiplier)
		minimunRatio := new(big.Float).SetFloat64(MinimumCollateralRatio)
		collateralRatio, _ := collateralRatioFloat.Float64()
		if collateralRatioFloat.Cmp(minimunRatio) == -1 {
			return sdkerrors.Wrapf(
				types.ErrTooLowCollateralRatio,
				fmt.Sprintf("collateral rate is too low. (%f%)", collateralRatio),
			)
		}
	}

	// CDP mint Mei coins
	err := keeper.SupplyKeeper.MintCoins(ctx, ModuleName, borrowAmountCoins)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMintCoin,
			"mint coin fail",
		)
	}

	// CDP send coins from module to sender
	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, msg.Sender, borrowAmountCoins)
	if err != nil {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"can't transfer coins from CDP module to sender",
		)
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	return nil
}
