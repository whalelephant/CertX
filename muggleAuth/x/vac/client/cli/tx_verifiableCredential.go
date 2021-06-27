package cli

import (
	"github.com/spf13/cobra"
	"strconv"
    "bufio"
    "fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	// "github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/client/utils"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
   	"github.com/cosmos/cosmos-sdk/crypto/hd"
 	bip39 "github.com/cosmos/go-bip39"
    "github.com/cosmos/cosmos-sdk/crypto/keyring"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdSendVerifiableCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-verifiableCredential [src-port] [src-channel] [verifier] [claim]",
		Short: "Send a verifiableCredential over IBC",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

            // TODO The subject should be generated, a new identifier which has a keypair stored in the keyring-backend
            // The verifier did should be the name of this new identifier so the user knows
			argsVerifier := string(args[2])
			argsClaim := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress().String()
        	var kr keyring.Keyring

            buf := bufio.NewReader(cmd.InOrStdin())
            backend := "test"
        	kr, err = keyring.New(sdk.KeyringServiceName(), backend, clientCtx.KeyringDir, buf)
        
        	if err != nil {
        		return err
        	}

            subject, err := createNewIdentifier(argsVerifier, kr)

            fmt.Println("subject: ", subject)

            // TODO  hard code the channel port with setting for the relayer
            // remove the need to input this
			srcPort := args[0]
			srcChannel := args[1]

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendVerifiableCredential(sender, srcPort, srcChannel, timeoutTimestamp, string(subject), string(argsVerifier), string(argsClaim))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
            return nil
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func createNewIdentifier(name string, kb keyring.Keyring) (subject string, e error) {
   	var err error

    // Warning: default override existing key
	_, err = kb.Key(name)
	if err == nil {
        err2 := kb.Delete(name)
        if err2 != nil {
            return "", err2
        }
	}


    // Create new key
    // We use the default algo to create new key
	keyringAlgos, _ := kb.SupportedAlgorithms()
    algoStr := "secp256k1"
	algo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
	if err != nil {
		return "", err
	}
    // Using Default values
	coinType := uint32(118)
	account := uint32(0)
	index := uint32(0) 
    hdPath := hd.CreateHDPath(coinType, account, index).String()

	// read entropy seed straight from tmcrypto.Rand and convert to mnemonic
    var mnemonicEntropySize = 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", err
	}

	//  Use entropy to generiic bip39 mnemonic
	var mnemonic, bip39Passphrase string
	mnemonic, err = bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

    // TODO: returned info will be needed for TX 
    // Add new key to the keyring backend (test)
    keyInfo, err := kb.NewAccount(name, mnemonic, bip39Passphrase, hdPath, algo)
	if err != nil {
		return "", err
	}
    addr := keyInfo.GetAddress()

	var subjectDid string
	subjectDid = "did:certX:"
    subjectDid += addr.String()[6:]

    return subjectDid, nil
}


