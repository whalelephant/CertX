package cli

import (
	"github.com/spf13/cobra"
	"strconv"
    "crypto/rand"
    "fmt"
    "encoding/base64"


	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/client/utils"
	"github.com/whalelephant/certX/MoM/x/employments/types"
)

var _ = strconv.Itoa(0)

func CmdSendECredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-eCredential [src-port] [src-channel] [subject] [claim]",
		Short: "Send a eCredential over IBC",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSubject := string(args[2])
			argsClaim := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

            // AES key 32 bytes
            // To be returned to the caller to give to required trusted entity
            key := make([]byte, 32)
            _, err = rand.Read(key[:])
            if err != nil {
                panic(err)
            }
            ekey := base64.StdEncoding.EncodeToString(key)
            fmt.Println(ekey)

			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendECredential(sender, srcPort, srcChannel, timeoutTimestamp, string(argsSubject), string(argsClaim), ekey)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
