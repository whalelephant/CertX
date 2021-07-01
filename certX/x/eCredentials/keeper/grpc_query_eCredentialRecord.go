package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/whalelephant/certX/certX/x/eCredentials/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ECredentialRecordAll(c context.Context, req *types.QueryAllECredentialRecordRequest) (*types.QueryAllECredentialRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var eCredentialRecords []*types.ECredentialRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	eCredentialRecordStore := prefix.NewStore(store, types.KeyPrefix(types.ECredentialRecordKey))

	pageRes, err := query.Paginate(eCredentialRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var eCredentialRecord types.ECredentialRecord
		if err := k.cdc.UnmarshalBinaryBare(value, &eCredentialRecord); err != nil {
			return err
		}

		eCredentialRecords = append(eCredentialRecords, &eCredentialRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllECredentialRecordResponse{ECredentialRecord: eCredentialRecords, Pagination: pageRes}, nil
}

func (k Keeper) ECredentialRecord(c context.Context, req *types.QueryGetECredentialRecordRequest) (*types.QueryGetECredentialRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var eCredentialRecord types.ECredentialRecord
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasECredentialRecord(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetECredentialRecordIDBytes(req.Id)), &eCredentialRecord)

	return &types.QueryGetECredentialRecordResponse{ECredentialRecord: &eCredentialRecord}, nil
}
