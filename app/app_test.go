package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestBCDExport(t *testing.T) {
	db := db.NewMemDB()
	mapp := NewMeichainApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, "")
	err := setGenesis(mapp)
	require.NoError(t, err)

	// Making a new app object with the db, so that initchain hasn't been called
	newBCapp := NewMeichainApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, "")
	_, _, err = newBCapp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := db.NewMemDB()
	mapp := NewMeichainApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, map[int64]bool{}, "")

	for acc := range maccPerms {
		require.True(t, mapp.bankKeeper.BlacklistedAddr(mapp.supplyKeeper.GetModuleAddress(acc)))
	}
}

func setGenesis(mapp *MeichainApp) error {
	genesisState := simapp.NewDefaultGenesisState()
	stateBytes, err := codec.MarshalJSONIndent(mapp.Codec(), genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	mapp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	mapp.Commit()
	return nil
}
