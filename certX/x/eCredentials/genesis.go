package eCredentials

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/certX/x/eCredentials/keeper"
	"github.com/whalelephant/certX/certX/x/eCredentials/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the eCredentialRecord
	for _, elem := range genState.ECredentialRecordList {
		k.SetECredentialRecord(ctx, *elem)
	}

	// Set eCredentialRecord count
	k.SetECredentialRecordCount(ctx, genState.ECredentialRecordCount)

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
	// Get all eCredentialRecord
	eCredentialRecordList := k.GetAllECredentialRecord(ctx)
	for _, elem := range eCredentialRecordList {
		elem := elem
		genesis.ECredentialRecordList = append(genesis.ECredentialRecordList, &elem)
	}

	// Set the current count
	genesis.ECredentialRecordCount = k.GetECredentialRecordCount(ctx)

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
