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

func (k Keeper) ProofAll(c context.Context, req *types.QueryAllProofRequest) (*types.QueryAllProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proofs []*types.Proof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proofStore := prefix.NewStore(store, types.KeyPrefix(types.ProofKey))

	pageRes, err := query.Paginate(proofStore, req.Pagination, func(key []byte, value []byte) error {
		var proof types.Proof
		if err := k.cdc.UnmarshalBinaryBare(value, &proof); err != nil {
			return err
		}

		proofs = append(proofs, &proof)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProofResponse{Proof: proofs, Pagination: pageRes}, nil
}

func (k Keeper) Proof(c context.Context, req *types.QueryGetProofRequest) (*types.QueryGetProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proof types.Proof
	ctx := sdk.UnwrapSDKContext(c)

    if !k.HasProof(ctx, req.Id) {
        return nil, sdkerrors.ErrKeyNotFound
    }

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetProofIDBytes(req.Id)), &proof)

	return &types.QueryGetProofResponse{Proof: &proof}, nil
}
