package gov

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermZero PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermSetPermissions PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermClaimValidator PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermClaimCouncilor PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermCreateSetPermissionsProposal PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermVoteSetPermissionProposal PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermUpsertTokenAlias PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermChangeTxFee PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermUpsertTokenRate PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermUpsertRole PermValue = 9
	// PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermUpsertDataRegistryProposal PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermVoteUpsertDataRegistryProposal PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermCreateSetNetworkPropertyProposal PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermVoteSetNetworkPropertyProposal PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermCreateUpsertTokenAliasProposal PermValue = 14
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermCreateSetPoorNetworkMessagesProposal PermValue = 15
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermVoteUpsertTokenAliasProposal PermValue = 16
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermCreateUpsertTokenRateProposal PermValue = 17
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermVoteUpsertTokenRateProposal PermValue = 18
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermVoteSetPoorNetworkMessagesProposal PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermCreateUnjailValidatorProposal PermValue = 20
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL",
	5:  "PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	15: "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
	16: "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	17: "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
	18: "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
	19: "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
	20: "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                    0,
	"PERMISSION_SET_PERMISSIONS":                         1,
	"PERMISSION_CLAIM_VALIDATOR":                         2,
	"PERMISSION_CLAIM_COUNCILOR":                         3,
	"PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL":         4,
	"PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL":           5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                      6,
	"PERMISSION_CHANGE_TX_FEE":                           7,
	"PERMISSION_UPSERT_TOKEN_RATE":                       8,
	"PERMISSION_UPSERT_ROLE":                             9,
	"PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL":           10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":      11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL":    12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":      13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":      14,
	"PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES":        15,
	"PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL":        16,
	"PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL":       17,
	"PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL":         18,
	"PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL": 19,
	"PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL":        20,
}
