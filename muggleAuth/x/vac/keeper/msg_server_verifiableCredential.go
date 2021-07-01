package keeper

import (
	"context"
	"fmt"
    "strconv"
    "strings"
    "encoding/base64"
    "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
)

func (k msgServer) SendVerifiableCredential(goCtx context.Context, msg *types.MsgSendVerifiableCredential) (*types.MsgSendVerifiableCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    fmt.Println("msg.Sender: ", msg.Sender);
    fmt.Println("msg.Claim: ", msg.Claim);

    claimId, err := strconv.Atoi(msg.Claim)
    if err != nil {
        fmt.Println("cannot convert claim: ", err);
        return nil, err
    } 

	// Checks that the element exists
	if !k.HasCredential(ctx, uint64(claimId)) {
        fmt.Println("No crednetial ");
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", claimId))
	}

	// Checks if the the msg sender is the same as the current owner
    claim := k.GetCredential(ctx, uint64(claimId))
    holder_addr := claim.GetHolder()[15:];
    if msg.Sender[7:] != holder_addr {
        fmt.Println("Wrong holder crednetial ");
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Holder")
    } 

    var externClaim string;
    switch numVac := claim.GetClaim(); numVac {
    case 0:
        externClaim = "None";
    case 1:
        externClaim = "Partial";
    case 2:
        externClaim = "Full"
    }

    credentialMsg := msg.Subject + msg.Verifier + claim.GetIssuer() + externClaim 

    // HACK WARNING: This keypath is the same as the client, demo purpose only! 
    // reader is not used since we are using test backend (not protected)
    keyringDir:= ".home"
    backend := "test"
    reader := strings.NewReader("")

    // The signer should be the creator of the credential
    signerAddr, err := sdk.AccAddressFromBech32(claim.GetCreator()) 
    kr, err := keyring.New(sdk.KeyringServiceName(), backend, keyringDir, reader)

    sig, _, err := kr.SignByAddress(signerAddr, []byte(credentialMsg))
    if err != nil {
        fmt.Println("Cannot sign");
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot sign")
    }

    // secp256k1 pubkey is 33 bytes
    // and sig is 64 bytes
    // or maybe base64.URLEncoding.EncodeToString(b)
    sigStr := base64.StdEncoding.EncodeToString(sig)

    // Construct the proof and store it
    var proof types.Proof;
    proof.Creator = claim.GetCreator()
    // Bech32 "cosmos" + "1" 
    proof.Holder = "did:MuugleAuth:" + msg.Sender[7:]
    proof.Issuer = claim.GetIssuer()
    proof.Verifier = msg.Verifier
    proof.Subject = msg.Subject
    proof.Signature = sigStr
	proof.Claim = externClaim 
    _ = k.AppendProof( ctx, proof)

	// Construct the packet
	var packet types.VerifiableCredentialPacketData
	packet.Subject = msg.Subject
	packet.Verifier = msg.Verifier
	packet.Issuer = claim.GetIssuer()
	packet.Claim = externClaim
    packet.Signature = sigStr
 
	// Transmit the packet
	e := k.TransmitVerifiableCredentialPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if e != nil {
		return nil, e
	}

	return &types.MsgSendVerifiableCredentialResponse{}, nil
}
