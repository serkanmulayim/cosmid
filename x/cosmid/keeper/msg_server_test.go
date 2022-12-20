package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/serkanmulayim/cosmid/testutil/keeper"
	"github.com/serkanmulayim/cosmid/x/cosmid/keeper"
	"github.com/serkanmulayim/cosmid/x/cosmid/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmidKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
