package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
)

func CmdCreateCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-credential [issuer] [holder] [claim]",
		Short: "Creates a new credential",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIssuer := string(args[0])
			argsHolder := string(args[1])
			argsClaim, _ := strconv.ParseInt(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCredential(clientCtx.GetFromAddress().String(), string(argsIssuer), string(argsHolder), int32(argsClaim))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-credential [id] [issuer] [holder] [claim]",
		Short: "Update a credential",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsIssuer := string(args[1])
			argsHolder := string(args[2])
			argsClaim, _ := strconv.ParseInt(args[3], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCredential(clientCtx.GetFromAddress().String(), id, string(argsIssuer), string(argsHolder), int32(argsClaim))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-credential [id] [issuer] [holder] [claim]",
		Short: "Delete a credential by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCredential(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
