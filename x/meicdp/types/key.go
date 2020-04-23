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
	CDPStoreKeyPrefix = []byte{0x01}

	// MsgCountStoreKey is a key for getting message count state variable
	MsgCountStoreKey = append([]byte(ModuleName), []byte("MsgCount")...)

	// MsgStoreKeyPrefix is a prefix for storing Message
	MsgStoreKeyPrefix = []byte{0x02}
)

// ResultStoreKey is a function to generate key for each result in store
// TODO: this is temporary function. Don't forget to remove after using this key
func ResultStoreKey(requestID oracle.RequestID) []byte {
	return append(ResultStoreKeyPrefix, int64ToBytes(int64(requestID))...)
}

// CDPStoreKey is a function to generate key for each CDP in store
func CDPStoreKey(account sdk.AccAddress) []byte {
	return append(CDPStoreKeyPrefix, []byte(account)...)
}

// MsgStoreKey is a function to generate key for each message in store
func MsgStoreKey(msgID uint64) []byte {
	return append(MsgStoreKeyPrefix, uint64ToBytes(msgID)...)
}

func int64ToBytes(num int64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(num))
	return result
}

func uint64ToBytes(num uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, num)
	return result
}
