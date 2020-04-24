package meicdp

import (
	"github.com/tharamalai/meichain/x/meicdp/keeper"
	"github.com/tharamalai/meichain/x/meicdp/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
	AtomUnit   = types.AtomUnit
	MeiUnit    = types.MeiUnit
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier
)

type (
	Keeper         = keeper.Keeper
	MsgRequestData = types.MsgRequestData
)
