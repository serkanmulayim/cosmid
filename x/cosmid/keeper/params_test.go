package keeper_test

import (
	"testing"

	testkeeper "github.com/serkanmulayim/cosmid/testutil/keeper"
	"github.com/serkanmulayim/cosmid/x/cosmid/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmidKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
