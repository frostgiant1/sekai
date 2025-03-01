package types

import (
	kiratypes "github.com/KiraCore/sekai/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Permissions: map[uint64]*Permissions{
			uint64(RoleSudo): NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimCouncilor,
				PermClaimValidator,
				PermCreateSetPermissionsProposal,
				PermVoteSetPermissionProposal,
				PermCreateSetNetworkPropertyProposal,
				PermVoteSetNetworkPropertyProposal,
				PermUpsertDataRegistryProposal,
				PermVoteUpsertDataRegistryProposal,
				PermCreateUpsertTokenAliasProposal,
				PermVoteUpsertTokenAliasProposal,
				PermCreateUpsertTokenRateProposal,
				PermVoteUpsertTokenRateProposal,
				PermUpsertRole,
				PermCreateUnjailValidatorProposal,
				PermCreateRoleProposal,
				PermVoteCreateRoleProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                    100,
			MaxTxFee:                    1000000,
			VoteQuorum:                  33,
			ProposalEndTime:             600, // 600 seconds / 10 mins
			ProposalEnactmentTime:       300, // 300 seconds / 5 mins
			EnableForeignFeePayments:    true,
			MischanceRankDecreaseAmount: 10,
			InactiveRankDecreasePercent: 50,      // 50%
			PoorNetworkMaxBankSend:      1000000, // 1M ukex
			MinValidators:               1,
			JailMaxTime:                 10, // 10 mins
			EnableTokenWhitelist:        false,
			EnableTokenBlacklist:        true,
		},
		ExecutionFees: []*ExecutionFee{
			{
				Name:              "Claim Validator Seat",
				TransactionType:   "claim-validator-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Governance Seat",
				TransactionType:   "claim-governance-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Proposal Type X",
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Vote Proposal Type X",
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Submit Proposal Type X",
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Veto Proposal Type X",
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Upsert Token Alias Execution Fee",
				TransactionType:   kiratypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Activate a validator",
				TransactionType:   kiratypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Pause a validator",
				TransactionType:   kiratypes.MsgTypePause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Unpause a validator",
				TransactionType:   kiratypes.MsgTypeUnpause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				kiratypes.MsgTypeProposalAssignPermission,
				kiratypes.MsgTypeProposalSetNetworkProperty,
				kiratypes.MsgTypeSetNetworkProperties,
				kiratypes.MsgTypeVoteProposal,
				kiratypes.MsgTypeClaimCouncilor,
				kiratypes.MsgTypeWhitelistPermissions,
				kiratypes.MsgTypeBlacklistPermissions,
				kiratypes.MsgTypeCreateRole,
				kiratypes.MsgTypeAssignRole,
				kiratypes.MsgTypeRemoveRole,
				kiratypes.MsgTypeWhitelistRolePermission,
				kiratypes.MsgTypeBlacklistRolePermission,
				kiratypes.MsgTypeRemoveWhitelistRolePermission,
				kiratypes.MsgTypeRemoveBlacklistRolePermission,
				kiratypes.MsgTypeClaimValidator,
				kiratypes.MsgTypeActivate,
				kiratypes.MsgTypePause,
				kiratypes.MsgTypeUnpause,
			},
		},
	}
}
