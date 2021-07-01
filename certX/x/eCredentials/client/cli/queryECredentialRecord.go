package cli

import (
	"strconv"
    "encoding/base64"
    "crypto/aes"
    "crypto/cipher"
	"context"
    "fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/whalelephant/certX/certX/x/eCredentials/types"
)

func CmdListECredentialRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-eCredentialRecord",
		Short: "list all eCredentialRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllECredentialRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ECredentialRecordAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowECredentialRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-eCredentialRecord [id] [key (optional)]",
		Short: "shows a eCredentialRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetECredentialRecordRequest{
				Id: id,
			}

			res, err := queryClient.ECredentialRecord(context.Background(), params)
			if err != nil {
				return err
			}


            if len(args[1]) > 0 {
                // on the other side
                record := res.ECredentialRecord.GetClaim()

                decodedClaim, err := base64.StdEncoding.DecodeString(record)
                key, err := base64.StdEncoding.DecodeString(args[1])

                c, err := aes.NewCipher(key)
                if err != nil {
                    fmt.Println(err)
                }
                gcmDecrypt, err := cipher.NewGCM(c)
                if err != nil {
                    fmt.Println(err)
                }
                nonceSize := gcmDecrypt.NonceSize()
                if len(decodedClaim) < nonceSize {
                    fmt.Println(err)
                }
                nonce, encryptedMessage := decodedClaim[:nonceSize], decodedClaim[nonceSize:]
                plaintext, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println("sig: ", string(plaintext[0:64]))
                fmt.Println("msg: ", string(plaintext[64:]))
            }
            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
