package meicdp

import (
	"github.com/tharamalai/meichain/x/meicdp/keeper"
	"github.com/tharamalai/meichain/x/meicdp/types"
)

const (
	ModuleName             = types.ModuleName
	RouterKey              = types.RouterKey
	StoreKey               = types.StoreKey
	CosmosHubChain         = types.CosmosHubChain
	BandChainID            = types.BandChainID
	OracleScriptID         = types.OracleScriptID
	AtomSymbol             = types.AtomSymbol
	MeiSymbol              = types.MeiSymbol
	AtomDecimal            = types.AtomDecimal
	MeiDecimal             = types.MeiDecimal
	AtomMultiplier         = types.AtomMultiplier
	AtomUnitPerAtom        = types.AtomUnitPerAtom
	MeiUnitPerMei          = types.MeiUnitPerMei
	MinimumCollateralRatio = types.MinimumCollateralRatio
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier
)

type (
	Keeper              = keeper.Keeper
	MsgLockCollateral   = types.MsgLockCollateral
	MsgUnlockCollateral = types.MsgUnlockCollateral
	MsgBorrowDebt       = types.MsgBorrowDebt
	MsgReturnDebt       = types.MsgReturnDebt
	MsgLiquidate        = types.MsgReturnDebt
	MsgSetSourceChannel = types.MsgSetSourceChannel
	DataRequest         = types.DataRequest
	Decoder             = types.Decoder
	Encoder             = types.Encoder
)
