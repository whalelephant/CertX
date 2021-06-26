package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/whalelephant/certX/certX/x/credential/types"
)

func createNRecvProof(keeper *Keeper, ctx sdk.Context, n int) []types.RecvProof {
	items := make([]types.RecvProof, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendRecvProof(ctx, items[i])
	}
	return items
}

func TestRecvProofGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecvProof(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetRecvProof(ctx, item.Id))
	}
}

func TestRecvProofExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecvProof(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasRecvProof(ctx, item.Id))
	}
}

func TestRecvProofRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecvProof(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRecvProof(ctx, item.Id)
		assert.False(t, keeper.HasRecvProof(ctx, item.Id))
	}
}

func TestRecvProofGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecvProof(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllRecvProof(ctx))
}

func TestRecvProofCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecvProof(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetRecvProofCount(ctx))
}
