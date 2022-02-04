package keeper_test

import (
	"testing"

	testkeeper "github.com/hashpays/hashchain/testutil/keeper"
	"github.com/hashpays/hashchain/x/hashchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HashchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
