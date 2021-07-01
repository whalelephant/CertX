package keeper

import (
	"github.com/whalelephant/certX/x/certx/types"
)

var _ types.QueryServer = Keeper{}
