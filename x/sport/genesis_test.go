package sport_test

import (
	"testing"

	keepertest "github.com/andrexbass/sport/testutil/keeper"
	"github.com/andrexbass/sport/testutil/nullify"
	"github.com/andrexbass/sport/x/sport"
	"github.com/andrexbass/sport/x/sport/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SportKeeper(t)
	sport.InitGenesis(ctx, *k, genesisState)
	got := sport.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
