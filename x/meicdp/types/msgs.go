package types

import (
	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgRequestData is a message for requesting a new data request to an existing oracle script.
type MsgRequestData struct {
	OracleScriptID oracle.OracleScriptID `json:"oracleScriptID"`
	SourceChannel  string                `json:"sourceChannel"`
	ClientID       string                `json:"clientID"`
	Calldata       []byte                `json:"calldata"`
	AskCount       int64                 `json:"askCount"`
	MinCount       int64                 `json:"minCount"`
	Sender         sdk.AccAddress        `json:"sender"`
}

// NewMsgRequestData creates a new MsgRequestData instance.
func NewMsgRequestData(
	oracleScriptID oracle.OracleScriptID,
	sourceChannel string,
	clientID string,
	calldata []byte,
	askCount int64,
	minCount int64,
	sender sdk.AccAddress,
) MsgRequestData {
	return MsgRequestData{
		OracleScriptID: oracleScriptID,
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		ClientID:       clientID,
		Sender:         sender,
	}
}

// Route implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Type() string { return "consuming" }

// ValidateBasic implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Sender address must not be empty.")
	}
	if msg.OracleScriptID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Oracle script id (%d) must be positive.", msg.OracleScriptID)
	}
	if msg.AskCount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Ask validator count (%d) must be positive.",
			msg.AskCount,
		)
	}
	if msg.AskCount < msg.MinCount {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Request validator count (%d) must not be less than minimum validator count (%d).",
			msg.AskCount,
			msg.MinCount,
		)
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgBorrowDebt is a message for borrow debt of Sender
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
func (msg MsgBorrowDebt) Type() string { return "borrow_debt" }

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
func (msg MsgLockCollateral) Type() string { return "lock_collateral" }

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

// MsgReturnDebt is a message for return debt of Sender
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
func (msg MsgReturnDebt) Type() string { return "return_debt" }

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
