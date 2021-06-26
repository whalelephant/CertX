package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
    "github.com/whalelephant/certX/muggleAuth/x/vac/types"
)

func createNProof(keeper *Keeper, ctx sdk.Context, n int) []types.Proof {
	items := make([]types.Proof, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendProof(ctx, items[i])
	}
	return items
}

func TestProofGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProof(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetProof(ctx, item.Id))
	}
}

func TestProofExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProof(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasProof(ctx, item.Id))
	}
}

func TestProofRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProof(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProof(ctx, item.Id)
		assert.False(t, keeper.HasProof(ctx, item.Id))
	}
}

func TestProofGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProof(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllProof(ctx))
}

func TestProofCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProof(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetProofCount(ctx))
}
