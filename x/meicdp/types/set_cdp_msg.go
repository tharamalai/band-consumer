package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetCDP is a message for unlocking collaterral of Sender
type MsgSetCDP struct {
	Sender sdk.AccAddress
}

// NewMsgSetCDP creates a new MsgSetCDP instance.
func NewMsgSetCDP(
	sender sdk.AccAddress,
) MsgSetCDP {
	return MsgSetCDP{
		Sender: sender,
	}
}

// Route implements the sdk.Msg interface for MsgSetCDP.
func (msg MsgSetCDP) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgSetCDP.
func (msg MsgSetCDP) Type() string { return "meicdp" }

// ValidateBasic implements the sdk.Msg interface for MsgSetCDP.
func (msg MsgSetCDP) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgSetCDP: Sender address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgSetCDP.
func (msg MsgSetCDP) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgSetCDP.
func (msg MsgSetCDP) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
