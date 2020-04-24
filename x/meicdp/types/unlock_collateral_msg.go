package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgUnlockCollateral is a message for unlocking collaterral of Sender
type MsgUnlockCollateral struct {
	Amount sdk.Coins
	Sender sdk.AccAddress
}

// NewMsgUnlockCollateral creates a new MsgUnlockCollateral instance.
func NewMsgUnlockCollateral(
	amount sdk.Coins,
	sender sdk.AccAddress,
) MsgUnlockCollateral {
	return MsgUnlockCollateral{
		Amount: amount,
		Sender: sender,
	}
}

// Route implements the sdk.Msg interface for MsgUnlockCollateral.
func (msg MsgUnlockCollateral) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgUnlockCollateral.
func (msg MsgUnlockCollateral) Type() string { return "unlock_collateral" }

// ValidateBasic implements the sdk.Msg interface for MsgUnlockCollateral.
func (msg MsgUnlockCollateral) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgUnlockCollateral: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgUnlockCollateral: Unlock amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgUnlockCollateral.
func (msg MsgUnlockCollateral) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgUnlockCollateral.
func (msg MsgUnlockCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
