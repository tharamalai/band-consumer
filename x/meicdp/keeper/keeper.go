package keeper

import (
	"bytes"
	"encoding/json"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.Marshaler
	BankKeeper    types.BankKeeper
	ChannelKeeper types.ChannelKeeper
}

// NewKeeper creates a new band consumer Keeper instance.
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, channelKeeper types.ChannelKeeper, bankKeeper types.BankKeeper) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		ChannelKeeper: channelKeeper,
		BankKeeper:    bankKeeper,
	}
}

func (k Keeper) SetResult(ctx sdk.Context, requestID oracle.RequestID, result []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ResultStoreKey(requestID), result)
}

func (k Keeper) GetResult(ctx sdk.Context, requestID oracle.RequestID) ([]byte, error) {
	if !k.HasResult(ctx, requestID) {
		return nil, sdkerrors.Wrapf(types.ErrItemNotFound,
			"GetResult: Result for request ID %d is not available.", requestID,
		)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.ResultStoreKey(requestID)), nil
}

func (k Keeper) HasResult(ctx sdk.Context, requestID oracle.RequestID) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ResultStoreKey(requestID))
}

//SetCDP - set CDP of user account to the store
func (k Keeper) SetCDP(ctx sdk.Context, account sdk.AccAddress, cdp types.CDP) {
	store := ctx.KVStore(k.storeKey)
	cdpBytes := new(bytes.Buffer)
	json.NewEncoder(cdpBytes).Encode(cdp)
	store.Set(types.CDPStoreKey(account), cdpBytes.Bytes())
}

//GetCDP - get CDP of user account from the store
func (k Keeper) GetCDP(ctx sdk.Context, account sdk.AccAddress) (types.CDP, error) {
	store := ctx.KVStore(k.storeKey)
	if k.HasCDP(ctx, account) {
		s := store.Get(types.CDPStoreKey(account))
		var cdp types.CDP
		if err := json.Unmarshal(s, &cdp); err != nil {
			return types.CDP{}, sdkerrors.Wrapf(types.ErrUnmarshalJSON,
				"GetCDP: CDP of %d is invalid json data.", account,
			)
		}

		return cdp, nil
	}
	return types.CDP{}, sdkerrors.Wrapf(types.ErrItemNotFound,
		"GetCDP: CDP of %d is not available.", account,
	)
}

// HasCDP - has CDP of this account on the store
func (k Keeper) HasCDP(ctx sdk.Context, account sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.CDPStoreKey(account))
}
