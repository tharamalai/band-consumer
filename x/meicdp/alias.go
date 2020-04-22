package consuming

import (
	"github.com/tharamalai/meichain/x/meicdp/keeper"
	"github.com/tharamalai/meichain/x/meicdp/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
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
