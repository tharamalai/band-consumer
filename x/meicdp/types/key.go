package types

import (
	"encoding/binary"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "meicdp"
	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// ResultStoreKeyPrefix is a prefix for storing result
	ResultStoreKeyPrefix = []byte{0xff}

	// CDPStoreKeyPrefix is a prefix for storing CDP
	CDPStoreKeyPrefix = []byte{0x00}
)

// ResultStoreKey is a function to generate key for each result in store
func ResultStoreKey(requestID oracle.RequestID) []byte {
	return append(ResultStoreKeyPrefix, int64ToBytes(int64(requestID))...)
}

// CDPStoreKey is a function to generate key for each CDP in store
func CDPStoreKey(account sdk.AccAddress) []byte {
	return append(CDPStoreKeyPrefix, []byte(account)...)
}

func int64ToBytes(num int64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(num))
	return result
}
