package keeper

import (
	"context"
    "strings"
    "fmt"
    "encoding/hex"
    "encoding/base64"
    "io"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "strconv"


	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/whalelephant/certX/MoM/x/employments/types"
)

func (k msgServer) SendECredential(goCtx context.Context, msg *types.MsgSendECredential) (*types.MsgSendECredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    fmt.Println("msg.Sender: ", msg.Sender);
    fmt.Println("msg.Claim: ", msg.Claim);

    claimId, err := strconv.Atoi(msg.Claim)
    if err != nil {
        fmt.Println("cannot convert claim: ", err);
        return nil, err
    } 

	// Checks that the element exists
	if !k.HasRecord(ctx, uint64(claimId)) {
        fmt.Println("No crednetial ");
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", claimId))
	}

	// Checks if the the msg sender is the same as the current owner
    record := k.GetRecord(ctx, uint64(claimId))
    // len: 13 - did:magician:
    holder_addr := record.GetSubject()[13:];
    fmt.Println("Holder: ", holder_addr)
    fmt.Println("Sender: ", msg.Sender[7:])
    if msg.Sender[7:] != holder_addr {
        fmt.Println("Wrong holder crednetial ");
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Holder")
    } 

    // create string from record 
    rec := record.GetRole()+ record.GetSince()
    recordStr := []byte(rec)

    // HACK WARNING: This keypath is the same as the client, demo purpose only! 
    // Reader is not used since we are using test backend (not protected)
    keyringDir:= ".home"
    backend := "test"
    reader := strings.NewReader("")

    // The signer should be the creator of the credential
    signerAddr, err := sdk.AccAddressFromBech32(record.GetCreator()) 
    kr, err := keyring.New(sdk.KeyringServiceName(), backend, keyringDir, reader)

    _, _, err = kr.SignByAddress(signerAddr, recordStr)
    if err != nil {
        fmt.Println("Cannot sign");
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot sign")
    }

    // encrypt with ekey  (sign and record)
    key, err := hex.DecodeString(msg.GetEkey())
    if err != nil {
        fmt.Println(err)
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Decoding error")
    }

    ciphr, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
        return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Error in eCredentials")
    }
    gcm, err := cipher.NewGCM(ciphr)
    if err != nil {
        fmt.Println(err)
        return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Error in eCredentials")
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println(err)
    }
    // text:= append(sig, recordStr...)
    text := recordStr
    claim := gcm.Seal(nonce, nonce, text, nil)
    claimStr := base64.StdEncoding.EncodeToString(claim)
    fmt.Println("claimStr: ", claimStr)

	// Construct the packet
	var packet types.ECredentialPacketData

	packet.Subject = msg.Subject
    // 20 works
    fmt.Println("claimStr len: ", len(claimStr))
    packet.Claim = claimStr 

	// Transmit the packet
	err = k.TransmitECredentialPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)

	if err != nil {
		return nil, err
	}

	return &types.MsgSendECredentialResponse{}, nil
}
