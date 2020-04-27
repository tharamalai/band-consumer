package types

import (
	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgBorrowDebt is a message for borrow debt of Sender
type MsgBorrowDebt struct {
	Amount uint64
	Sender sdk.AccAddress
}

// NewMsgBorrowDebt creates a new MsgBorrowDebt instance.
func NewMsgBorrowDebt(
	amount uint64,
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
	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBorrowDebt: Borrow amount must more than 0.")
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
	Amount uint64
	Sender sdk.AccAddress
}

// NewMsgLockCollateral creates a new MsgLockCollateral instance.
func NewMsgLockCollateral(
	amount uint64,
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
	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgLockCollateral: Lock amount must more than 0.")
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
	Amount uint64
	Sender sdk.AccAddress
}

// NewMsgReturnDebt creates a new MsgReturnDebt instance.
func NewMsgReturnDebt(
	amount uint64,
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
	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReturnDebt: Return amount must more than 0.")
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
	Amount uint64
	Sender sdk.AccAddress
}

// NewMsgUnlockCollateral creates a new MsgUnlockCollateral instance.
func NewMsgUnlockCollateral(
	amount uint64,
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
	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgUnlockCollateral: Unlock amount must more than 0.")
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

// MsgSetSoruceChannel is a message for setting source channel to other chain
type MsgSetSourceChannel struct {
	ChainName     string         `json:"chain_name"`
	SourcePort    string         `json:"source_port"`
	SourceChannel string         `json:"source_channel"`
	Signer        sdk.AccAddress `json:"signer"`
}

func NewMsgSetSourceChannel(
	chainName, sourcePort, sourceChannel string,
	signer sdk.AccAddress,
) MsgSetSourceChannel {
	return MsgSetSourceChannel{
		ChainName:     chainName,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Signer:        signer,
	}
}

// Route implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Type() string { return "set_channel" }

// ValidateBasic implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) ValidateBasic() error {
	// TODO: Add validate basic
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// GetSignBytes implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// DataRequest is a message for requesting a new data request to an existing oracle script.
type DataRequest struct {
	OracleScriptID oracle.OracleScriptID `json:"oracleScriptID"`
	SourceChannel  string                `json:"sourceChannel"`
	ChainID        string                `json:"chainID"`
	Port           string                `json:"port"`
	ClientID       string                `json:"clientID"`
	Calldata       []byte                `json:"calldata"`
	AskCount       int64                 `json:"askCount"`
	MinCount       int64                 `json:"minCount"`
	Sender         sdk.AccAddress        `json:"sender"`
}

// NewDataRequest creates a new DataRequest instance.
func NewDataRequest(
	oracleScriptID oracle.OracleScriptID,
	sourceChannel string,
	chainID string,
	port string,
	clientID string,
	calldata []byte,
	askCount int64,
	minCount int64,
	sender sdk.AccAddress,
) DataRequest {
	return DataRequest{
		OracleScriptID: oracleScriptID,
		SourceChannel:  sourceChannel,
		ChainID:        chainID,
		Port:           port,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		ClientID:       clientID,
		Sender:         sender,
	}
}
