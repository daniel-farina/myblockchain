package myblockchain_test

import (
	"testing"

	keepertest "github.com/daniel-farina/myblockchain/testutil/keeper"
	"github.com/daniel-farina/myblockchain/x/myblockchain"
	"github.com/daniel-farina/myblockchain/x/myblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyblockchainKeeper(t)
	myblockchain.InitGenesis(ctx, *k, genesisState)
	got := myblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
