package keeper

import (
	"github.com/whalelephant/certX/certX/x/credentials/types"
)

var _ types.QueryServer = Keeper{}
