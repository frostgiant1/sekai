package testslashing

import (
	"github.com/KiraCore/sekai/x/slashing/types"
)

// TestParams construct default slashing params for tests.
// Have to change these parameters for tests
// lest the tests take forever
func TestParams() types.Params {
	params := types.DefaultParams()
	params.SignedBlocksWindow = 1000
	params.DowntimeInactiveDuration = 60 * 60

	return params
}
