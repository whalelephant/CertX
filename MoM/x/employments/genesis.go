package employments

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/MoM/x/employments/keeper"
	"github.com/whalelephant/certX/MoM/x/employments/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the record
	for _, elem := range genState.RecordList {
		k.SetRecord(ctx, *elem)
	}

	// Set record count
	k.SetRecordCount(ctx, genState.RecordCount)

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all record
	recordList := k.GetAllRecord(ctx)
	for _, elem := range recordList {
		elem := elem
		genesis.RecordList = append(genesis.RecordList, &elem)
	}

	// Set the current count
	genesis.RecordCount = k.GetRecordCount(ctx)

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
