package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgReturnDebt is a message for unlocking collaterral of Sender
type MsgReturnDebt struct {
	Amount sdk.Coins
	Sender sdk.AccAddress
}

// NewMsgReturnDebt creates a new MsgReturnDebt instance.
func NewMsgReturnDebt(
	amount sdk.Coins,
	sender sdk.AccAddress,
) MsgReturnDebt {
	return MsgReturnDebt{
		Amount: amount,
		Sender: sender,
	}
}

// Route implements the sdk.Msg interface for MsgReturnDebt.
func (msg MsgReturnDebt) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgReturnDebt.
func (msg MsgReturnDebt) Type() string { return "meicdp" }

// ValidateBasic implements the sdk.Msg interface for MsgReturnDebt.
func (msg MsgReturnDebt) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReturnDebt: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReturnDebt: Return amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgReturnDebt.
func (msg MsgReturnDebt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgReturnDebt.
func (msg MsgReturnDebt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
