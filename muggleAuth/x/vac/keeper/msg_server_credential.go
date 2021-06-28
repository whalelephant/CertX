package keeper

import (
	"context"
    "strings"
	"fmt"

    "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
)

func (k msgServer) CreateCredential(goCtx context.Context, msg *types.MsgCreateCredential) (*types.MsgCreateCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // store the standard pubkey as the identifier secp256k1 
    // in real did, we will have this in the did doc
	var issuerDid string
	issuerDid = "did:muggleAuth:"
    creatorAdd, err := sdk.AccAddressFromBech32(msg.Creator) 
    if err != nil {
        return nil, err
    }

    // HACK WARNING: This keypath is the same as the client, demo purpose only! 
    // reader is not used since we are using test backend (not protected)
    keyringDir:= ".home"
    backend := "test"
    reader := strings.NewReader("")
    kr, err := keyring.New(sdk.KeyringServiceName(), backend, keyringDir, reader)
    keyInfo, err := kr.KeyByAddress(creatorAdd)
    issuerDid += keyInfo.GetPubKey().Address().String()

    fmt.Println("issuerDid: ", issuerDid)

	id := k.AppendCredential(
		ctx,
		msg.Creator,
		issuerDid,
		msg.Holder,
		msg.Claim,
	)

	return &types.MsgCreateCredentialResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCredential(goCtx context.Context, msg *types.MsgUpdateCredential) (*types.MsgUpdateCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var credential = types.Credential{
		Creator: msg.Creator,
		Id:      msg.Id,
		Issuer:  msg.Issuer,
		Holder:  msg.Holder,
		Claim:   msg.Claim,
	}

	// Checks that the element exists
	if !k.HasCredential(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCredentialOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCredential(ctx, credential)

	return &types.MsgUpdateCredentialResponse{}, nil
}

func (k msgServer) DeleteCredential(goCtx context.Context, msg *types.MsgDeleteCredential) (*types.MsgDeleteCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasCredential(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetCredentialOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCredential(ctx, msg.Id)

	return &types.MsgDeleteCredentialResponse{}, nil
}
