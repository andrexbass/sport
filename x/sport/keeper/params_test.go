package keeper_test

import (
	"testing"

	testkeeper "github.com/andrexbass/sport/testutil/keeper"
	"github.com/andrexbass/sport/x/sport/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SportKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
