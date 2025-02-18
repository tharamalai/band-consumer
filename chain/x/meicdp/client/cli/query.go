package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	meiCdpCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the meicdp module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	meiCdpCmd.AddCommand(flags.GetCommands(
		GetCmdReadCDP(storeKey, cdc),
	)...)

	return meiCdpCmd
}

// GetCmdReadCDP - Get CDP by user account
func GetCmdReadCDP(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "cdp",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/cdp/%s", queryRoute, account),
				nil,
			)

			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(string(res))
		},
	}
}
