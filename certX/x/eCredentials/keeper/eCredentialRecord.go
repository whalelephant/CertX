package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/certX/x/eCredentials/types"
	"strconv"
)

// GetECredentialRecordCount get the total number of eCredentialRecord
func (k Keeper) GetECredentialRecordCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordCountKey))
	byteKey := types.KeyPrefix(types.ECredentialRecordCountKey)
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

// SetECredentialRecordCount set the total number of eCredentialRecord
func (k Keeper) SetECredentialRecordCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordCountKey))
	byteKey := types.KeyPrefix(types.ECredentialRecordCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendECredentialRecord appends a eCredentialRecord in the store with a new id and update the count
func (k Keeper) AppendECredentialRecord(
	ctx sdk.Context,
	eCredentialRecord types.ECredentialRecord,
) uint64 {
	// Create the eCredentialRecord
	count := k.GetECredentialRecordCount(ctx)

	// Set the ID of the appended value
	eCredentialRecord.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&eCredentialRecord)
	store.Set(GetECredentialRecordIDBytes(eCredentialRecord.Id), appendedValue)

	// Update eCredentialRecord count
	k.SetECredentialRecordCount(ctx, count+1)

	return count
}

// SetECredentialRecord set a specific eCredentialRecord in the store
func (k Keeper) SetECredentialRecord(ctx sdk.Context, eCredentialRecord types.ECredentialRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	b := k.cdc.MustMarshalBinaryBare(&eCredentialRecord)
	store.Set(GetECredentialRecordIDBytes(eCredentialRecord.Id), b)
}

// GetECredentialRecord returns a eCredentialRecord from its id
func (k Keeper) GetECredentialRecord(ctx sdk.Context, id uint64) types.ECredentialRecord {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	var eCredentialRecord types.ECredentialRecord
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetECredentialRecordIDBytes(id)), &eCredentialRecord)
	return eCredentialRecord
}

// HasECredentialRecord checks if the eCredentialRecord exists in the store
func (k Keeper) HasECredentialRecord(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	return store.Has(GetECredentialRecordIDBytes(id))
}

// GetECredentialRecordOwner returns the creator of the eCredentialRecord
func (k Keeper) GetECredentialRecordOwner(ctx sdk.Context, id uint64) string {
	return k.GetECredentialRecord(ctx, id).Creator
}

// RemoveECredentialRecord removes a eCredentialRecord from the store
func (k Keeper) RemoveECredentialRecord(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	store.Delete(GetECredentialRecordIDBytes(id))
}

// GetAllECredentialRecord returns all eCredentialRecord
func (k Keeper) GetAllECredentialRecord(ctx sdk.Context) (list []types.ECredentialRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ECredentialRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ECredentialRecord
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetECredentialRecordIDBytes returns the byte representation of the ID
func GetECredentialRecordIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetECredentialRecordIDFromBytes returns ID in uint64 format from a byte array
func GetECredentialRecordIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
