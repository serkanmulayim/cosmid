package cosmid_test

import (
	"testing"

	keepertest "github.com/serkanmulayim/cosmid/testutil/keeper"
	"github.com/serkanmulayim/cosmid/testutil/nullify"
	"github.com/serkanmulayim/cosmid/x/cosmid"
	"github.com/serkanmulayim/cosmid/x/cosmid/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmidKeeper(t)
	cosmid.InitGenesis(ctx, *k, genesisState)
	got := cosmid.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
