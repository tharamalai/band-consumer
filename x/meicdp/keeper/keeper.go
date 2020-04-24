package keeper

import (
	"encoding/binary"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	BankKeeper    types.BankKeeper
	ChannelKeeper types.ChannelKeeper
}

// NewKeeper creates a new Mei CDP Keeper instance.
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, channelKeeper types.ChannelKeeper, bankKeeper types.BankKeeper) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		ChannelKeeper: channelKeeper,
		BankKeeper:    bankKeeper,
	}
}

// TODO: Temporary keeper function for testing. Don't forget to remove.
func (k Keeper) SetResult(ctx sdk.Context, requestID oracle.RequestID, result []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ResultStoreKey(requestID), result)
}

// TODO: Temporary keeper function for testing. Don't forget to remove.
func (k Keeper) GetResult(ctx sdk.Context, requestID oracle.RequestID) ([]byte, error) {
	if !k.HasResult(ctx, requestID) {
		return nil, sdkerrors.Wrapf(types.ErrItemNotFound,
			"GetResult: Result for request ID %d is not available.", requestID,
		)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.ResultStoreKey(requestID)), nil
}

// TODO: Temporary keeper function for testing. Don't forget to remove.
func (k Keeper) HasResult(ctx sdk.Context, requestID oracle.RequestID) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ResultStoreKey(requestID))
}

// SetCDP - set CDP into the store
func (k Keeper) SetCDP(ctx sdk.Context, cdp types.CDP) {
	store := ctx.KVStore(k.storeKey)
	c := k.cdc.MustMarshalBinaryBare(cdp)
	store.Set(types.CDPStoreKey(cdp.Owner), c)
}

//GetCDP - get CDP of user account from the store
func (k Keeper) GetCDP(ctx sdk.Context, account sdk.AccAddress) types.CDP {
	store := ctx.KVStore(k.storeKey)
	if k.HasCDP(ctx, account) {
		c := store.Get(types.CDPStoreKey(account))
		var cdp types.CDP
		k.cdc.MustUnmarshalBinaryBare(c, &cdp)
		return cdp
	}

	atomToken := sdk.NewCoin("uatom", sdk.NewInt(0))
	collateralCoins := sdk.NewCoins(atomToken)

	meiToken := sdk.NewCoin("umei", sdk.NewInt(0))
	debtCoins := sdk.NewCoins(meiToken)

	return types.NewCDP(collateralCoins, debtCoins, account)
}

// HasCDP - has CDP of this account on the store
func (k Keeper) HasCDP(ctx sdk.Context, account sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.CDPStoreKey(account))
}

//SetMsg - set Msg of this message ID to the store
func (k Keeper) SetMsg(ctx sdk.Context, msgID uint64, msg sdk.Msg) {
	store := ctx.KVStore(k.storeKey)
	m := k.cdc.MustMarshalBinaryBare(msg)
	store.Set(types.MsgStoreKey(msgID), m)
}

//GetMsg - get Msg by msgID from the store
func (k Keeper) GetMsg(ctx sdk.Context, msgID uint64) (sdk.Msg, error) {
	store := ctx.KVStore(k.storeKey)
	if k.HasMsg(ctx, msgID) {
		m := store.Get(types.MsgStoreKey(msgID))
		var msg sdk.Msg
		k.cdc.MustUnmarshalBinaryBare(m, &msg)
		return msg, nil
	}
	return nil, sdkerrors.Wrapf(types.ErrItemNotFound,
		"GetMsg: Msg of %d is not available.", msgID,
	)
}

// HasMsg - has Msg of this msgID on the store
func (k Keeper) HasMsg(ctx sdk.Context, msgID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.MsgStoreKey(msgID))
}

// GetMsgCount returns a number of all messages on the store
func (k Keeper) GetMsgCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.MsgCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

//GetNextMsgCount returns and increments a number of the next msg
func (k Keeper) GetNextMsgCount(ctx sdk.Context) uint64 {
	msgCount := k.GetMsgCount(ctx)
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(msgCount + 1)
	store.Set(types.MsgCountStoreKey, bz)
	return msgCount + 1
}
