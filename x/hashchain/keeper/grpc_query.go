package keeper

import (
	"github.com/hashpays/hashchain/x/hashchain/types"
)

var _ types.QueryServer = Keeper{}
