package keeper

import (
	"github.com/serkanmulayim/cosmid/x/cosmid/types"
)

var _ types.QueryServer = Keeper{}
