package keeper

import (
	"github.com/whalelephant/certX/x/credentials/types"
)

var _ types.QueryServer = Keeper{}
