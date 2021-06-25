package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CredentialAll(c context.Context, req *types.QueryAllCredentialRequest) (*types.QueryAllCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var credentials []*types.Credential
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	credentialStore := prefix.NewStore(store, types.KeyPrefix(types.CredentialKey))

	pageRes, err := query.Paginate(credentialStore, req.Pagination, func(key []byte, value []byte) error {
		var credential types.Credential
		if err := k.cdc.UnmarshalBinaryBare(value, &credential); err != nil {
			return err
		}

		credentials = append(credentials, &credential)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCredentialResponse{Credential: credentials, Pagination: pageRes}, nil
}

func (k Keeper) Credential(c context.Context, req *types.QueryGetCredentialRequest) (*types.QueryGetCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var credential types.Credential
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCredential(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCredentialIDBytes(req.Id)), &credential)

	return &types.QueryGetCredentialResponse{Credential: &credential}, nil
}
