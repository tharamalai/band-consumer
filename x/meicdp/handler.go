package meicdp

import (
	"fmt"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
				fmt.Println("I GOT DATA", responseData.Result, responseData.ResolveTime)
				// handleOraclePacket(ctx, keeper, responseData)
				return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal oracle packet data")

		case types.MsgLockCollateral:
			return handleMsgLockCollateral(ctx, keeper, msg)
		case types.MsgReturnDebt:
			return handleMsgReturnDebt(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgLockCollateral(ctx sdk.Context, keeper Keeper, msg types.MsgLockCollateral) (*sdk.Result, error) {

	cdp := keeper.GetCDP(ctx, msg.Sender)

	// Accumulate collateral on CDP
	lockAmount := sdk.NewCoin(types.AtomUnit, sdk.NewInt(int64(msg.Amount)))
	lockAmountCoins := sdk.NewCoins(lockAmount)

	//  new collateral
	cdp.CollateralAmount = cdp.CollateralAmount + msg.Amount

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Transfer collateral to the module account. Transaction fails if sender's balance is insufficient.
	moduleAddress := types.GetMeiCDPAddress()
	err := keeper.BankKeeper.SendCoins(ctx, msg.Sender, moduleAddress, lockAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient fund")
	}

	return &sdk.Result{}, nil
}

func handleMsgReturnDebt(ctx sdk.Context, keeper Keeper, msg types.MsgReturnDebt) (*sdk.Result, error) {

	cdp := keeper.GetCDP(ctx, msg.Sender)

	// Subtract debt on CDP
	returnAmount := sdk.NewCoin(types.AtomUnit, sdk.NewInt(int64(msg.Amount)))
	returnAmountCoins := sdk.NewCoins(returnAmount)

	// new debt
	cdp.DebtAmount = cdp.DebtAmount - msg.Amount

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// Transfer Mei from user to CDP. Transaction fails if sender's balance is insufficient.
	moduleAddress := types.GetMeiCDPAddress()
	err := keeper.BankKeeper.SendCoins(ctx, msg.Sender, moduleAddress, returnAmountCoins)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}
