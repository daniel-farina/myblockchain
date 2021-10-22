package keeper

import (
	"github.com/daniel-farina/myblockchain/x/myblockchain/types"
)

var _ types.QueryServer = Keeper{}
