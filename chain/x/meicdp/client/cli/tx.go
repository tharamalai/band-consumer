package cli

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

const (
	flagName     = "name"
	flagCalldata = "calldata"
	flagClientID = "client-id"
	flagAskCount = "ask-count"
	flagMinCount = "min-count"
	flagChannel  = "channel"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	meiCdpCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "meicdp transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	meiCdpCmd.AddCommand(flags.PostCommands(
		GetCmdLockCollateral(cdc),
		GetCmdUnlockCollateral(cdc),
		GetCmdReturnDebt(cdc),
		GetCmdBorrowDebt(cdc),
		GetCmdLiquidate(cdc),
		GetCmdSetChannel(cdc),
		GetCmdMagic(cdc),
	)...)

	return meiCdpCmd
}

// GetCmdLockCollateral implements lock collateral handler
func GetCmdLockCollateral(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock [amount]",
		Short: "Lock collateral to the CDP.",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Lock collateral to the CDP.
Example:
$ %s tx meicdp lock 100000
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgLockCollateral(
				amount,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdUnlockCollateral implements unlock collateral handler
func GetCmdUnlockCollateral(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock [amount]",
		Short: "Unlock collateral from the CDP.",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Unlock collateral from the CDP.
Example:
$ %s tx meicdp unlock 100000
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnlockCollateral(
				amount,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdReturnDebt implements return debt handler
func GetCmdReturnDebt(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "return [amount]",
		Short: "Return debt to the CDP.",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Return debt to the CDP.
Example:
$ %s tx meicdp return 100000
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgReturnDebt(
				amount,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdBorrowDebt implements return debt handler
func GetCmdBorrowDebt(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow [amount]",
		Short: "Borrow debt from the CDP.",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Borrow debt from the CDP.
Example:
$ %s tx meicdp borrow 100000
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrowDebt(
				amount,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdLiquidate implements liquidate CDP handler
func GetCmdLiquidate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate [cdp]",
		Short: "Liquidate the CDP.",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Liquidate the CDP.
Example:
$ %s tx meicdp liquidate cosmos1rdajkxwtw4fz9c9u044z7qzn9t6q4eqn0dzxhk
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			cdpOwner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return errors.Unwrap(fmt.Errorf("invalid cdp owner address"))
			}

			msg := types.NewMsgLiquidate(
				cdpOwner,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdSetChannel implements the set channel command handler.
func GetCmdSetChannel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-channel [chain-id] [port] [channel-id]",
		Short: "Register a verified channel",
		Args:  cobra.ExactArgs(3),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Register a verified channel.
Example:
$ %s tx meicdp set-cahnnel bandchain meicdp dbdfgsdfsd
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			msg := types.NewMsgSetSourceChannel(
				args[0],
				args[1],
				args[2],
				cliCtx.GetFromAddress(),
			)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

//TODO: remove this. testing function.
func GetCmdMagic(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "magic [cdp] [liquidator]",
		Short: "Add debt to the CDP.",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Add debt to the CDP.
Example:
$ %s tx meicdp magic cosmos1rdajkxwtw4fz9c9u044z7qzn9t6q4eqn0dzxhk cosmos12ysrelpaf8566aehgtf4j096qatkqsf26f707u
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			cdpOwner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return errors.Unwrap(fmt.Errorf("invalid cdp owner address"))
			}

			liquidator, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return errors.Unwrap(fmt.Errorf("invalid liquidate address"))
			}

			msg := types.NewMsgAddDebtAdmin(
				cdpOwner,
				liquidator,
				cliCtx.GetFromAddress(),
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
