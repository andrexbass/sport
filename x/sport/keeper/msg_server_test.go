package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/andrexbass/sport/testutil/keeper"
	"github.com/andrexbass/sport/x/sport/keeper"
	"github.com/andrexbass/sport/x/sport/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SportKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
