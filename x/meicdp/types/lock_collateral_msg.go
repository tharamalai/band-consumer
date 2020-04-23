package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgLockCollateral is a message for lock collaterral of Sender
type MsgLockCollateral struct {
	Amount sdk.Coins
	Sender sdk.AccAddress
}

// NewMsgLockCollateral creates a new MsgLockCollateral instance.
func NewMsgLockCollateral(
	amount sdk.Coins,
	sender sdk.AccAddress,
) MsgLockCollateral {
	return MsgLockCollateral{
		Amount: amount,
		Sender: sender,
	}
}

// Route implements the sdk.Msg interface for MsgLockCollateral.
func (msg MsgLockCollateral) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgLockCollateral.
func (msg MsgLockCollateral) Type() string { return "meicdp" }

// ValidateBasic implements the sdk.Msg interface for MsgLockCollateral.
func (msg MsgLockCollateral) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgLockCollateral: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgLockCollateral: Lock amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgLockCollateral.
func (msg MsgLockCollateral) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgLockCollateral.
func (msg MsgLockCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
