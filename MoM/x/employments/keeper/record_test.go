package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/whalelephant/certX/MoM/x/employments/types"
)

func createNRecord(keeper *Keeper, ctx sdk.Context, n int) []types.Record {
	items := make([]types.Record, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendRecord(ctx, items[i])
	}
	return items
}

func TestRecordGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetRecord(ctx, item.Id))
	}
}

func TestRecordExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasRecord(ctx, item.Id))
	}
}

func TestRecordRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRecord(ctx, item.Id)
		assert.False(t, keeper.HasRecord(ctx, item.Id))
	}
}

func TestRecordGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllRecord(ctx))
}

func TestRecordCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetRecordCount(ctx))
}
