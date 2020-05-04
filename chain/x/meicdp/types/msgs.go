package types

import (
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

// MsgLiquidate is a message for liquidate CDP
type MsgLiquidate struct {
	CdpOwner   sdk.AccAddress
	Liquidator sdk.AccAddress
}

// NewMsgLiquidate creates a new MsgLiquidate instance.
func NewMsgLiquidate(
	cdpOwner sdk.AccAddress,
	liquidator sdk.AccAddress,
) MsgLiquidate {
	return MsgLiquidate{
		CdpOwner:   cdpOwner,
		Liquidator: liquidator,
	}
}

// Route implements the sdk.Msg interface for MsgLiquidate.
func (msg MsgLiquidate) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgLiquidate.
func (msg MsgLiquidate) Type() string { return "liquidate" }

// ValidateBasic implements the sdk.Msg interface for MsgLiquidate.
func (msg MsgLiquidate) ValidateBasic() error {
	if msg.CdpOwner.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgLiquidate: CdpOwner address must not be empty.")
	}
	if msg.Liquidator.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgLiquidate: Liquidator address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgLiquidate.
func (msg MsgLiquidate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Liquidator}
}

// GetSignBytes implements the sdk.Msg interface for MsgLiquidate.
func (msg MsgLiquidate) GetSignBytes() []byte {
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

type MsgAddDebtAdmin struct {
	CdpOwner   sdk.AccAddress `json:"cdpOwner"`
	Liquidator sdk.AccAddress `json:"liquidator"`
	Admin      sdk.AccAddress `json:"admin"`
}

func NewMsgAddDebtAdmin(
	cdpOwner sdk.AccAddress,
	liquidator sdk.AccAddress,
	admin sdk.AccAddress,
) MsgAddDebtAdmin {
	return MsgAddDebtAdmin{
		CdpOwner:   cdpOwner,
		Liquidator: liquidator,
		Admin:      admin,
	}
}

// Route implements the sdk.Msg interface for MsgAddDebtAdmin.
func (msg MsgAddDebtAdmin) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgAddDebtAdmin.
func (msg MsgAddDebtAdmin) Type() string { return "set_debt" }

// ValidateBasic implements the sdk.Msg interface for MsgAddDebtAdmin.
func (msg MsgAddDebtAdmin) ValidateBasic() error {
	// TODO: Add validate basic
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgAddDebtAdmin.
func (msg MsgAddDebtAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Admin}
}

// GetSignBytes implements the sdk.Msg interface for MsgAddDebtAdmin.
func (msg MsgAddDebtAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
