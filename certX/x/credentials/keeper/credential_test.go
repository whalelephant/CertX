package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/whalelephant/certX/x/credentials/types"
)

func createNCredential(keeper *Keeper, ctx sdk.Context, n int) []types.Credential {
	items := make([]types.Credential, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendCredential(ctx, items[i])
	}
	return items
}

func TestCredentialGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCredential(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCredential(ctx, item.Id))
	}
}

func TestCredentialExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCredential(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCredential(ctx, item.Id))
	}
}

func TestCredentialRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCredential(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCredential(ctx, item.Id)
		assert.False(t, keeper.HasCredential(ctx, item.Id))
	}
}

func TestCredentialGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCredential(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCredential(ctx))
}

func TestCredentialCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCredential(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCredentialCount(ctx))
}
