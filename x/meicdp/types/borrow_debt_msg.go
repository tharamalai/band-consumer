package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgBorrowDebt is a message for unlocking collaterral of Sender
type MsgBorrowDebt struct {
	Amount sdk.Coins
	Sender sdk.AccAddress
}

// NewMsgBorrowDebt creates a new MsgBorrowDebt instance.
func NewMsgBorrowDebt(
	amount sdk.Coins,
	sender sdk.AccAddress,
) MsgBorrowDebt {
	return MsgBorrowDebt{
		Amount: amount,
		Sender: sender,
	}
}

// Route implements the sdk.Msg interface for MsgBorrowDebt.
func (msg MsgBorrowDebt) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgBorrowDebt.
func (msg MsgBorrowDebt) Type() string { return "meicdp" }

// ValidateBasic implements the sdk.Msg interface for MsgBorrowDebt.
func (msg MsgBorrowDebt) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBorrowDebt: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBorrowDebt: Borrow amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBorrowDebt.
func (msg MsgBorrowDebt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgBorrowDebt.
func (msg MsgBorrowDebt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
