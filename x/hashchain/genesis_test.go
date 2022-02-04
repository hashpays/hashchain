package hashchain_test

import (
	"testing"

	keepertest "github.com/hashpays/hashchain/testutil/keeper"
	"github.com/hashpays/hashchain/testutil/nullify"
	"github.com/hashpays/hashchain/x/hashchain"
	"github.com/hashpays/hashchain/x/hashchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HashchainKeeper(t)
	hashchain.InitGenesis(ctx, *k, genesisState)
	got := hashchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
