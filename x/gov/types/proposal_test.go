package types

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewProposal_ProposalIsPendingByDefault(t *testing.T) {
	proposal, err := NewProposal(
		1234,
		NewAssignPermissionProposal(
			types.AccAddress{0x12},
			PermSetPermissions,
		),
		time.Now(),
		time.Now(),
		time.Now(),
	)

	require.NoError(t, err)
	require.Equal(t, Pending, proposal.Result)
}
