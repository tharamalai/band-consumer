package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidBasicMsg        = sdkerrors.Register(ModuleName, 1, "InvalidBasicMsg")
	ErrBadDataValue           = sdkerrors.Register(ModuleName, 2, "BadDataValue")
	ErrUnauthorizedPermission = sdkerrors.Register(ModuleName, 3, "UnauthorizedPermission")
	ErrItemDuplication        = sdkerrors.Register(ModuleName, 4, "ItemDuplication")
	ErrItemNotFound           = sdkerrors.Register(ModuleName, 5, "ItemNotFound")
	ErrInvalidState           = sdkerrors.Register(ModuleName, 6, "InvalidState")
	ErrBadWasmExecution       = sdkerrors.Register(ModuleName, 7, "BadWasmExecution")
	ErrUnmarshalJSON          = sdkerrors.Register(ModuleName, 8, "InvalidCDPJSONData")
	ErrMintCoin               = sdkerrors.Register(ModuleName, 9, "BurnMintFail")
	ErrBurnCoin               = sdkerrors.Register(ModuleName, 10, "BurnCoinFail")
	ErrInvalidChannel         = sdkerrors.Register(ModuleName, 11, "InvalidChannel")
	ErrUnknownClientID        = sdkerrors.Register(ModuleName, 12, "UnknownClientID")
	ErrTooLowCollateralRatio  = sdkerrors.Register(ModuleName, 13, "TooLowCollateralRatio")
	ErrMsgNotFound            = sdkerrors.Register(ModuleName, 14, "MsgNotFound")
	ErrRequestOracleData      = sdkerrors.Register(ModuleName, 15, "ErrRequestOracleData")
	ErrResponseOracleData     = sdkerrors.Register(ModuleName, 16, "ErrResponseOracleData")
)
