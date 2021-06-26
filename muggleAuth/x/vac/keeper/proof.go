package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetProofCount get the total number of proof
func (k Keeper) GetProofCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofCountKey))
	byteKey := types.KeyPrefix(types.ProofCountKey)
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

// SetProofCount set the total number of proof
func (k Keeper) SetProofCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofCountKey))
	byteKey := types.KeyPrefix(types.ProofCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendProof appends a proof in the store with a new id and update the count
func (k Keeper) AppendProof(
    ctx sdk.Context,
    proof types.Proof,
) uint64 {
	// Create the proof
    count := k.GetProofCount(ctx)

    // Set the ID of the appended value
    proof.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
    appendedValue := k.cdc.MustMarshalBinaryBare(&proof)
    store.Set(GetProofIDBytes(proof.Id), appendedValue)

    // Update proof count
    k.SetProofCount(ctx, count+1)

    return count
}

// SetProof set a specific proof in the store
func (k Keeper) SetProof(ctx sdk.Context, proof types.Proof) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	b := k.cdc.MustMarshalBinaryBare(&proof)
	store.Set(GetProofIDBytes(proof.Id), b)
}

// GetProof returns a proof from its id
func (k Keeper) GetProof(ctx sdk.Context, id uint64) types.Proof {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	var proof types.Proof
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetProofIDBytes(id)), &proof)
	return proof
}

// HasProof checks if the proof exists in the store
func (k Keeper) HasProof(ctx sdk.Context, id uint64) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	return store.Has(GetProofIDBytes(id))
}

// GetProofOwner returns the creator of the proof
func (k Keeper) GetProofOwner(ctx sdk.Context, id uint64) string {
    return k.GetProof(ctx, id).Creator
}

// RemoveProof removes a proof from the store
func (k Keeper) RemoveProof(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	store.Delete(GetProofIDBytes(id))
}

// GetAllProof returns all proof
func (k Keeper) GetAllProof(ctx sdk.Context) (list []types.Proof) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProofKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proof
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetProofIDBytes returns the byte representation of the ID
func GetProofIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetProofIDFromBytes returns ID in uint64 format from a byte array
func GetProofIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
