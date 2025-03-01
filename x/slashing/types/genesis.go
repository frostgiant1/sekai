package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, signingInfos []SigningInfo, missedBlocks []ValidatorMissedBlocks,
) *GenesisState {

	return &GenesisState{
		Params:       params,
		SigningInfos: signingInfos,
		MissedBlocks: missedBlocks,
	}
}

// NewMissedBlock creates a new MissedBlock instance
func NewMissedBlock(index int64, missed bool) MissedBlock {
	return MissedBlock{
		Index:  index,
		Missed: missed,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:       DefaultParams(),
		SigningInfos: []SigningInfo{},
		MissedBlocks: []ValidatorMissedBlocks{},
	}
}

// ValidateGenesis validates the slashing genesis parameters
func ValidateGenesis(data GenesisState) error {
	minSign := data.Params.MinSignedPerWindow
	if minSign.IsNegative() || minSign.GT(sdk.OneDec()) {
		return fmt.Errorf("min signed per window should be less than or equal to one and greater than zero, is %s", minSign.String())
	}

	downtimeInactive := data.Params.DowntimeInactiveDuration
	if downtimeInactive < 1*time.Minute {
		return fmt.Errorf("downtime unblond duration must be at least 1 minute, is %s", downtimeInactive.String())
	}

	signedWindow := data.Params.SignedBlocksWindow
	if signedWindow < 10 {
		return fmt.Errorf("signed blocks window must be at least 10, is %d", signedWindow)
	}

	return nil
}
