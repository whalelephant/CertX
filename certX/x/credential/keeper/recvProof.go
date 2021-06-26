package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/certX/x/credential/types"
	"strconv"
)

// GetRecvProofCount get the total number of recvProof
func (k Keeper) GetRecvProofCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofCountKey))
	byteKey := types.KeyPrefix(types.RecvProofCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetRecvProofCount set the total number of recvProof
func (k Keeper) SetRecvProofCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofCountKey))
	byteKey := types.KeyPrefix(types.RecvProofCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendRecvProof appends a recvProof in the store with a new id and update the count
func (k Keeper) AppendRecvProof(
	ctx sdk.Context,
	recvProof types.RecvProof,
) uint64 {
	// Create the recvProof
	count := k.GetRecvProofCount(ctx)

	// Set the ID of the appended value
	recvProof.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&recvProof)
	store.Set(GetRecvProofIDBytes(recvProof.Id), appendedValue)

	// Update recvProof count
	k.SetRecvProofCount(ctx, count+1)

	return count
}

// SetRecvProof set a specific recvProof in the store
func (k Keeper) SetRecvProof(ctx sdk.Context, recvProof types.RecvProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	b := k.cdc.MustMarshalBinaryBare(&recvProof)
	store.Set(GetRecvProofIDBytes(recvProof.Id), b)
}

// GetRecvProof returns a recvProof from its id
func (k Keeper) GetRecvProof(ctx sdk.Context, id uint64) types.RecvProof {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	var recvProof types.RecvProof
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRecvProofIDBytes(id)), &recvProof)
	return recvProof
}

// HasRecvProof checks if the recvProof exists in the store
func (k Keeper) HasRecvProof(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	return store.Has(GetRecvProofIDBytes(id))
}

// GetRecvProofOwner returns the creator of the recvProof
func (k Keeper) GetRecvProofOwner(ctx sdk.Context, id uint64) string {
	return k.GetRecvProof(ctx, id).Creator
}

// RemoveRecvProof removes a recvProof from the store
func (k Keeper) RemoveRecvProof(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	store.Delete(GetRecvProofIDBytes(id))
}

// GetAllRecvProof returns all recvProof
func (k Keeper) GetAllRecvProof(ctx sdk.Context) (list []types.RecvProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecvProofKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RecvProof
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRecvProofIDBytes returns the byte representation of the ID
func GetRecvProofIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRecvProofIDFromBytes returns ID in uint64 format from a byte array
func GetRecvProofIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
