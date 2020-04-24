package meicdp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
		case MsgRequestData:
			sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, "meichain", msg.SourceChannel)
			if !found {
				return nil, sdkerrors.Wrapf(
					sdkerrors.ErrUnknownRequest,
					"unknown channel %s port meichain",
					msg.SourceChannel,
				)
			}
			destinationPort := sourceChannelEnd.Counterparty.PortID
			destinationChannel := sourceChannelEnd.Counterparty.ChannelID
			sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
				ctx, "meichain", msg.SourceChannel,
			)
			if !found {
				return nil, sdkerrors.Wrapf(
					sdkerrors.ErrUnknownRequest,
					"unknown sequence number for channel %s port oracle",
					msg.SourceChannel,
				)
			}
			packet := oracle.NewOracleRequestPacketData(
				msg.ClientID, msg.OracleScriptID, hex.EncodeToString(msg.Calldata),
				msg.AskCount, msg.MinCount,
			)
			err := keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
				sequence, "meichain", msg.SourceChannel, destinationPort, destinationChannel,
				1000000000, // Arbitrarily high timeout for now
			))
			if err != nil {
				return nil, err
			}
			return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

		case types.MsgSetSourceChannel:
			return handleSetSourceChannel(ctx, msg, keeper)

		case channeltypes.MsgPacket:
			var responseDataPacket oracle.OracleResponsePacketData
			if err := types.ModuleCdc.UnmarshalJSON(msg.GetData(), &responseDataPacket); err == nil {
				fmt.Println("I GOT DATA", responseDataPacket.Result, responseDataPacket.ResolveTime)
				handleOracleRespondPacketData(ctx, keeper, responseDataPacket)
				return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal oracle packet data")

		case types.MsgLockCollateral:
			return handleMsgLockCollateral(ctx, keeper, msg)

		case types.MsgReturnDebt:
			return handleMsgReturnDebt(ctx, keeper, msg)

		case types.MsgUnlockCollateral:
			msgCount := keeper.GetMsgCount(ctx)

			// setup oracle request
			bandChainID := "bandchain"
			port := "meicdp"
			oracleScriptID := oracle.OracleScriptID(3)
			clientID := fmt.Sprintf("Msg:%d", msgCount)
			calldata := make([]byte, 8)
			binary.LittleEndian.PutUint64(calldata, 1000000)
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
				msg.Sender,
			)

			// Set message to the store for waiting the oracle response packet.
			keeper.SetMsg(ctx, msgCount, msg)

			err = requestOracle(ctx, keeper, dataRequest)
			if err != nil {
				return nil, err
			}

			return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

		case types.MsgBorrowDebt:
			msgCount := keeper.GetMsgCount(ctx)

			// setup oracle request
			bandChainID := "bandchain"
			port := "meicdp"
			oracleScriptID := oracle.OracleScriptID(3)
			clientID := fmt.Sprintf("Msg:%d", msgCount)
			calldata := make([]byte, 8)
			binary.LittleEndian.PutUint64(calldata, 1000000)
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
				msg.Sender,
			)

			// Set message to the store for waiting the oracle response packet.
			keeper.SetMsg(ctx, msgCount, msg)

			err = requestOracle(ctx, keeper, dataRequest)
			if err != nil {
				return nil, err
			}

			return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgLockCollateral(ctx sdk.Context, keeper Keeper, msg types.MsgLockCollateral) (*sdk.Result, error) {

	cdp, err := keeper.GetCDP(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	// Accumulate collateral on CDP
	newCollateral := cdp.CollateralAmount.Add(msg.Amount...)
	fmt.Println("newCollateral", newCollateral)
	cdp.CollateralAmount = newCollateral

	fmt.Println("cdp", cdp)

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Transfer collateral to the module account. Transaction fails if sender's balance is insufficient.
	moduleAddress := types.GetMeiCDPAddress()
	err = keeper.BankKeeper.SendCoins(ctx, msg.Sender, moduleAddress, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient fund")
	}

	return &sdk.Result{}, nil
}

func handleMsgReturnDebt(ctx sdk.Context, keeper Keeper, msg types.MsgReturnDebt) (*sdk.Result, error) {

	cdp, err := keeper.GetCDP(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	// Subtract debt on CDP
	newDebt := cdp.DebtAmount.Sub(msg.Amount)
	fmt.Println("newCollateral", newDebt)
	cdp.DebtAmount = newDebt

	fmt.Println("cdp", cdp)

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Transfer Mei from user to CDP. Transaction fails if sender's balance is insufficient.
	moduleAddress := types.GetMeiCDPAddress()
	err = keeper.BankKeeper.SendCoins(ctx, msg.Sender, moduleAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

func handleSetSourceChannel(ctx sdk.Context, msg types.MsgSetSourceChannel, keeper Keeper) (*sdk.Result, error) {
	keeper.SetChannel(ctx, msg.ChainName, msg.SourcePort, msg.SourceChannel)
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func requestOracle(ctx sdk.Context, keeper Keeper, dataReq types.DataRequest) error {

	channelID, err := keeper.GetChannel(ctx, dataReq.ChainID, dataReq.Port)

	sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, "meichain", channelID)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port meichain",
			channelID,
		)
	}

	destinationPort := sourceChannelEnd.Counterparty.PortID
	destinationChannel := sourceChannelEnd.Counterparty.ChannelID
	sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
		ctx, "meichain", channelID,
	)

	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			channelID,
		)
	}

	packet := oracle.NewOracleRequestPacketData(
		dataReq.ClientID, dataReq.OracleScriptID, string(dataReq.Calldata),
		dataReq.AskCount, dataReq.MinCount,
	)

	err = keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
		sequence, "meichain", channelID, destinationPort, destinationChannel,
		1000000000, // Arbitrarily high timeout for now
	))

	if err != nil {
		return err
	}

	return nil
}

func handleOracleRespondPacketData(ctx sdk.Context, keeper Keeper, packet oracle.OracleResponsePacketData) (*sdk.Result, error) {
	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return nil, sdkerrors.Wrapf(types.ErrUnknownClientID, "unknown client id %s", packet.ClientID)
	}

	msgID, err := strconv.ParseUint(clientID[1], 10, 64)
	if err != nil {
		return nil, err
	}

	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return nil, err
	}

	result, err := types.DecodeResult(rawResult)
	if err != nil {
		return nil, err
	}

	msg, err := keeper.GetMsg(ctx, msgID)
	if err != nil {
		return nil, err
	}

	switch msg := msg.(type) {
	case types.MsgUnlockCollateral:
		err := handleMsgUnlockCollatearl(ctx, keeper, msg, result)
		if err != nil {
			return nil, err
		}

	case types.MsgBorrowDebt:
		err := handleMsgBorrowDebt(ctx, keeper, msg, result)
		if err != nil {
			return nil, err
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

}

// handleMsgUnlockCollatearl handles the unlock collateral message after receives oracle packet
func handleMsgUnlockCollatearl(ctx sdk.Context, keeper Keeper, msg types.MsgUnlockCollateral, packetResult types.Result) error {
	cdp, err := keeper.GetCDP(ctx, msg.Sender)
	if err != nil {
		return err
	}

	// TODO: Calculate new collateral ratio

	newCollateral := cdp.CollateralAmount.Sub(msg.Amount)
	fmt.Println("newCollateral", newCollateral)
	cdp.CollateralAmount = newCollateral

	fmt.Println("cdp", cdp)

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Move collateral from CDP module to sender account
	moduleAddress := types.GetMeiCDPAddress()
	err = keeper.BankKeeper.SendCoins(ctx, moduleAddress, msg.Sender, msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient fund")
	}

	return nil
}

// handleMsgBorrowDebt handles the borrow debt message after receives oracle packet
func handleMsgBorrowDebt(ctx sdk.Context, keeper Keeper, msg types.MsgBorrowDebt, packetResult types.Result) error {
	cdp, err := keeper.GetCDP(ctx, msg.Sender)
	if err != nil {
		return err
	}

	// TODO: Calculate new collateral ratio

	newDebt := cdp.DebtAmount.Add(msg.Amount...)
	fmt.Println("newDebt", newDebt)
	cdp.DebtAmount = newDebt

	fmt.Println("cdp", cdp)

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Move debt from CDP module to sender account
	moduleAddress := types.GetMeiCDPAddress()
	err = keeper.BankKeeper.SendCoins(ctx, moduleAddress, msg.Sender, msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient fund")
	}

	return nil
}
