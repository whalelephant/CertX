package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/whalelephant/certX/certX/x/credential/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RecvProofAll(c context.Context, req *types.QueryAllRecvProofRequest) (*types.QueryAllRecvProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var recvProofs []*types.RecvProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	recvProofStore := prefix.NewStore(store, types.KeyPrefix(types.RecvProofKey))

	pageRes, err := query.Paginate(recvProofStore, req.Pagination, func(key []byte, value []byte) error {
		var recvProof types.RecvProof
		if err := k.cdc.UnmarshalBinaryBare(value, &recvProof); err != nil {
			return err
		}

		recvProofs = append(recvProofs, &recvProof)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRecvProofResponse{RecvProof: recvProofs, Pagination: pageRes}, nil
}

func (k Keeper) RecvProof(c context.Context, req *types.QueryGetRecvProofRequest) (*types.QueryGetRecvProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var recvProof types.RecvProof
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasRecvProof(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRecvProofIDBytes(req.Id)), &recvProof)

	return &types.QueryGetRecvProofResponse{RecvProof: &recvProof}, nil
}
