package types

import (
	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DataRequest is a message for requesting a new data request to an existing oracle script.
type DataRequest struct {
	OracleScriptID oracle.OracleScriptID `json:"oracleScriptID"`
	SourceChannel  string                `json:"sourceChannel"`
	ChainID        string                `json:"chainID"`
	Port           string                `json:"port"`
	ClientID       string                `json:"clientID"`
	Calldata       []byte                `json:"calldata"`
	AskCount       int64                 `json:"askCount"`
	MinCount       int64                 `json:"minCount"`
	Sender         sdk.AccAddress        `json:"sender"`
}

// NewDataRequest creates a new DataRequest instance.
func NewDataRequest(
	oracleScriptID oracle.OracleScriptID,
	sourceChannel string,
	chainID string,
	port string,
	clientID string,
	calldata []byte,
	askCount int64,
	minCount int64,
	sender sdk.AccAddress,
) DataRequest {
	return DataRequest{
		OracleScriptID: oracleScriptID,
		SourceChannel:  sourceChannel,
		ChainID:        chainID,
		Port:           port,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		ClientID:       clientID,
		Sender:         sender,
	}
}
