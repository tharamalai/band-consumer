package meicdp

import (
	"fmt"
	"math/big"

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

		case MsgLockCollateral:
			return handleMsgLockCollateral(ctx, keeper, msg)
		case MsgReturnDebt:
			return handleMsgReturnDebt(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgLockCollateral(ctx sdk.Context, keeper Keeper, msg MsgLockCollateral) (*sdk.Result, error) {

	cdp := keeper.GetCDP(ctx, msg.Sender)

	lockAmount := sdk.NewCoin(types.AtomUnit, sdk.NewInt(int64(msg.Amount)))
	lockAmountCoins := sdk.NewCoins(lockAmount)

	//  Accumulate collateral on CDP
	lockAmountInt := new(big.Int).SetUint64(msg.Amount)
	collateralAmountInt := new(big.Int).SetUint64(cdp.CollateralAmount)
	collateralAmountInt.Add(collateralAmountInt, lockAmountInt)
	if !collateralAmountInt.IsUint64() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidBasicMsg, "invalid lock amount. collateral must more than or equals 0.")
	}

	cdp.CollateralAmount = collateralAmountInt.Uint64()

	// Transfer collateral to the module account. Transaction fails if sender's balance is insufficient.
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, lockAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "can't transfer tokens from sender to CDP")
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
		return nil, sdkerrors.Wrapf(types.ErrInvalidBasicMsg, "invalid return amount. debt must more than or equals 0.")
	}
	cdp.DebtAmount = debtAmountInt.Uint64()

	// TODO: Pay fee

	// Transfer Mei from user to CDP. Transaction fails if sender's balance is insufficient.
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "can't transfer tokens from sender to CDP")
	}

	// Store CDP
	keeper.SetCDP(ctx, cdp)

	// CDP burn returned coins
	err = keeper.SupplyKeeper.BurnCoins(ctx, ModuleName, returnAmountCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrBurnCoin, "burn coin fail")
	}

	return &sdk.Result{}, nil
}
