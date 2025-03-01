package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewValidatorSigningInfo creates a new ValidatorSigningInfo instance
//nolint:interfacer
func NewValidatorSigningInfo(
	condAddr sdk.ConsAddress, startHeight, indexOffset int64,
	inactivatedUntil time.Time, tombstoned bool, missedBlocksCounter int64,
) ValidatorSigningInfo {

	return ValidatorSigningInfo{
		Address:             condAddr.String(),
		StartHeight:         startHeight,
		IndexOffset:         indexOffset,
		InactiveUntil:       inactivatedUntil,
		Tombstoned:          tombstoned,
		MissedBlocksCounter: missedBlocksCounter,
	}
}

// String implements the stringer interface for ValidatorSigningInfo
func (i ValidatorSigningInfo) String() string {
	return fmt.Sprintf(`Validator Signing Info:
  Address:               %s
  Start Height:          %d
  Index Offset:          %d
  Inactivated Until:          %v
  Tombstoned:            %t
  Missed Blocks Counter: %d`,
		i.Address, i.StartHeight, i.IndexOffset, i.InactiveUntil,
		i.Tombstoned, i.MissedBlocksCounter)
}

// unmarshal a validator signing info from a store value
func UnmarshalValSigningInfo(cdc codec.Marshaler, value []byte) (signingInfo ValidatorSigningInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &signingInfo)
	return signingInfo, err
}
