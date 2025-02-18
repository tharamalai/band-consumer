package types

import (
	"encoding/binary"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "meicdp"

	// RouterKey is the module name router key
	RouterKey = ModuleName

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// AtomUnit is unit of ATOM
	AtomUnit = "uatom"

	// MeiUnit is unit of MEI
	MeiUnit = "umei"

	// CosmosHubChain Cosmos Hub chain ID
	CosmosHubChain = "band-cosmoshub"

	// AtomSymbol is symbol of ATOM
	AtomSymbol = "ATOM"

	// MeiSymbol is symbol of MEI
	MeiSymbol = "MEI"

	// AtomDecimal is decimal of ATOM
	AtomDecimal = 6

	// MeiDecimal is decimal of MEI
	MeiDecimal = 6

	// AtomUnitPerAtom is amount of atom unit per one atom (10^6)
	AtomUnitPerAtom = 1000000

	// MeiUnitPerMei is amount of mei unit per one mei (10^6)
	MeiUnitPerMei = 1000000

	// AtomMultiplier is multiplier of Atom price request
	AtomMultiplier = 1000000

	// MinimumCollateralRatio is 150%
	MinimumCollateralRatio = 150

	//BandChainID - Bandchain ID
	BandChainID = "ibc-bandchain"

	//OracleScriptID is crypto price script (Borsh version)
	OracleScriptID = 2
)

var (
	// ResultStoreKeyPrefix is a prefix for storing result
	// TODO: this is temporary prefix. Don't forget to remove
	ResultStoreKeyPrefix = []byte{0xff}

	// CDPStoreKeyPrefix is a prefix for storing CDP
	CDPStoreKeyPrefix = []byte{0x01}

	// MsgCountStoreKey is a key for getting message count state variable
	MsgCountStoreKey = append([]byte(ModuleName), []byte("MsgCount")...)

	// MsgStoreKeyPrefix is a prefix for storing Message
	MsgStoreKeyPrefix = []byte{0x02}

	// ChannelStoreKeyPrefix is a prefix for storing channel
	ChannelStoreKeyPrefix = []byte{0x03}
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

// ChannelStoreKey is a function to generate key for each verified channel in store
func ChannelStoreKey(chainName, channelPort string) []byte {
	buf := append(ChannelStoreKeyPrefix, []byte(chainName)...)
	buf = append(buf, []byte(channelPort)...)
	return buf
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
