package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/x/credentials/types"
	"strconv"
)

// GetCredentialCount get the total number of credential
func (k Keeper) GetCredentialCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialCountKey))
	byteKey := types.KeyPrefix(types.CredentialCountKey)
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

// SetCredentialCount set the total number of credential
func (k Keeper) SetCredentialCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialCountKey))
	byteKey := types.KeyPrefix(types.CredentialCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCredential appends a credential in the store with a new id and update the count
func (k Keeper) AppendCredential(
	ctx sdk.Context,
	credential types.Credential,
) uint64 {
	// Create the credential
	count := k.GetCredentialCount(ctx)

	// Set the ID of the appended value
	credential.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&credential)
	store.Set(GetCredentialIDBytes(credential.Id), appendedValue)

	// Update credential count
	k.SetCredentialCount(ctx, count+1)

	return count
}

// SetCredential set a specific credential in the store
func (k Keeper) SetCredential(ctx sdk.Context, credential types.Credential) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	b := k.cdc.MustMarshalBinaryBare(&credential)
	store.Set(GetCredentialIDBytes(credential.Id), b)
}

// GetCredential returns a credential from its id
func (k Keeper) GetCredential(ctx sdk.Context, id uint64) types.Credential {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	var credential types.Credential
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCredentialIDBytes(id)), &credential)
	return credential
}

// HasCredential checks if the credential exists in the store
func (k Keeper) HasCredential(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	return store.Has(GetCredentialIDBytes(id))
}

// GetCredentialOwner returns the creator of the credential
func (k Keeper) GetCredentialOwner(ctx sdk.Context, id uint64) string {
	return k.GetCredential(ctx, id).Creator
}

// RemoveCredential removes a credential from the store
func (k Keeper) RemoveCredential(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	store.Delete(GetCredentialIDBytes(id))
}

// GetAllCredential returns all credential
func (k Keeper) GetAllCredential(ctx sdk.Context) (list []types.Credential) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredentialKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Credential
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCredentialIDBytes returns the byte representation of the ID
func GetCredentialIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCredentialIDFromBytes returns ID in uint64 format from a byte array
func GetCredentialIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
