package keeper

import (
	"fmt"
	"strconv"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryResult = "result"
	QueryCDP    = "cdp"
)

// NewQuerier is the module level router for state queries.
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryResult:
			return queryResult(ctx, path[1:], req, keeper)
		case QueryCDP:
			return queryCDP(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
		}
	}
}

// queryResult is a query function to get result by request ID.
func queryResult(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) ([]byte, error) {
	if len(path) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "must specify the request id")
	}
	id, err := strconv.ParseInt(path[0], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, fmt.Sprintf("wrong format for requestid %s", err.Error()))
	}
	return keeper.GetResult(ctx, oracle.RequestID(id))
}

// queryCDP is a query function to get CDP by account.
func queryCDP(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) ([]byte, error) {
	if len(path) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "must specify the account")
	}

	accAccount, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid account format")
	}

	cdp, err := keeper.GetCDP(ctx, accAccount)
	if err != nil {
		return nil, err
	}

	return keeper.cdc.MustMarshalJSON(cdp), nil
}
