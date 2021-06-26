package credential

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/whalelephant/certX/certX/x/credential/keeper"
	"github.com/whalelephant/certX/certX/x/credential/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the recvProof
	for _, elem := range genState.RecvProofList {
		k.SetRecvProof(ctx, *elem)
	}

	// Set recvProof count
	k.SetRecvProofCount(ctx, genState.RecvProofCount)

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
	// Get all recvProof
	recvProofList := k.GetAllRecvProof(ctx)
	for _, elem := range recvProofList {
		elem := elem
		genesis.RecvProofList = append(genesis.RecvProofList, &elem)
	}

	// Set the current count
	genesis.RecvProofCount = k.GetRecvProofCount(ctx)

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
