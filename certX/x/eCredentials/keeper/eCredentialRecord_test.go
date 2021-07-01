package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/whalelephant/certX/certX/x/eCredentials/types"
)

func createNECredentialRecord(keeper *Keeper, ctx sdk.Context, n int) []types.ECredentialRecord {
	items := make([]types.ECredentialRecord, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendECredentialRecord(ctx, items[i])
	}
	return items
}

func TestECredentialRecordGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNECredentialRecord(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetECredentialRecord(ctx, item.Id))
	}
}

func TestECredentialRecordExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNECredentialRecord(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasECredentialRecord(ctx, item.Id))
	}
}

func TestECredentialRecordRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNECredentialRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveECredentialRecord(ctx, item.Id)
		assert.False(t, keeper.HasECredentialRecord(ctx, item.Id))
	}
}

func TestECredentialRecordGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNECredentialRecord(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllECredentialRecord(ctx))
}

func TestECredentialRecordCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNECredentialRecord(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetECredentialRecordCount(ctx))
}
