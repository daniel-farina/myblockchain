package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/daniel-farina/myblockchain/testutil/keeper"
	"github.com/daniel-farina/myblockchain/x/myblockchain/keeper"
	"github.com/daniel-farina/myblockchain/x/myblockchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
