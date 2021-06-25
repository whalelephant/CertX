package vac

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/muggleAuth/x/vac/keeper"
	"github.com/whalelephant/certX/muggleAuth/x/vac/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the credential
	for _, elem := range genState.CredentialList {
		k.SetCredential(ctx, *elem)
	}

	// Set credential count
	k.SetCredentialCount(ctx, uint64(len(genState.CredentialList)))

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
	// Get all credential
	credentialList := k.GetAllCredential(ctx)
	for _, elem := range credentialList {
		elem := elem
		genesis.CredentialList = append(genesis.CredentialList, &elem)
	}

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
