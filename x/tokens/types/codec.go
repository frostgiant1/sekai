package types

import (
	functionmeta "github.com/KiraCore/sekai/function_meta"
	govtypes "github.com/KiraCore/sekai/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec register codec and metadata
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpsertTokenAlias{}, "kiraHub/MsgUpsertTokenAlias", nil)
	functionmeta.AddNewFunction((&MsgUpsertTokenAlias{}).Type(), `{
		"description": "MsgUpsertTokenAlias represents a message to register token alias.",
		"parameters": {
			"symbol": {
				"type":        "string",
				"description": "Ticker (eg. ATOM, KEX, BTC)."
			},
			"name": {
				"type":        "string",
				"description": "Token Name (e.g. Cosmos, Kira, Bitcoin)."
			},
			"icon": {
				"type":        "string",
				"description": "Graphical Symbol (url link to graphics)."
			},
			"decimals": {
				"type":        "uint32",
				"description": "Integer number of max decimals."
			},
			"denoms": {
				"type":        "array<string>",
				"description": "An array of token denoms to be aliased."
			},
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgUpsertTokenRate{}, "kiraHub/MsgUpsertTokenRate", nil)
	functionmeta.AddNewFunction((&MsgUpsertTokenRate{}).Type(), `{
		"description": "MsgUpsertTokenRate represents a message to register token rate.",
		"parameters": {
			"denom": {
				"type":        "string",
				"description": "denomination target."
			},
			"rate": {
				"type":        "float",
				"description": "Exchange rate in terms of KEX token. e.g. 0.1, 10.5"
			},
			"fee_payments": {
				"type":        "bool",
				"description": "defining if it is enabled or disabled as fee payment method."
			},
			"proposer": {
				"type":        "address",
				"description": "proposer who propose this message."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgProposalUpsertTokenAlias{}, "kiraHub/MsgProposalUpsertTokenAlias", nil)
	functionmeta.AddNewFunction((&MsgProposalUpsertTokenAlias{}).Type(), `{
		"description": "MsgProposalUpsertTokenAlias defines a proposal message to upsert token alias.",
		"parameters": {
			"symbol": {
				"type":        "string",
				"description": "Ticker (eg. ATOM, KEX, BTC)."
			},
			"name": {
				"type":        "string",
				"description": "Token Name (e.g. Cosmos, Kira, Bitcoin)."
			},
			"icon": {
				"type":        "string",
				"description": "Graphical Symbol (url link to graphics)."
			},
			"decimals": {
				"type":        "uint32",
				"description": "Integer number of max decimals."
			},
			"denoms": {
				"type":        "array<string>",
				"description": "An array of token denoms to be aliased."
			},
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgProposalUpsertTokenRates{}, "kiraHub/MsgProposalUpsertTokenRates", nil)
	functionmeta.AddNewFunction((&MsgProposalUpsertTokenRates{}).Type(), `{
		"description": "MsgProposalUpsertTokenRates defines a proposal message to upsert token rate.",
		"parameters": {
			"denom": {
				"type":        "string",
				"description": "denominator, ukex, uatom etc."
			},
			"rate": {
				"type":        "decimal",
				"description": "e.g. 1.20"
			},
			"fee_payments": {
				"type":        "bool",
				"description": "describe if specified denom can be used for fee payment."
			},
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgProposalTokensWhiteBlackChange{}, "kiraHub/MsgProposalTokensWhiteBlackChange", nil)
	functionmeta.AddNewFunction((&MsgProposalTokensWhiteBlackChange{}).Type(), `{
		"description": "MsgProposalTokensWhiteBlackChange defines a proposal message to change tokens blacklists / whitelists.",
		"parameters": {
			"is_blacklist": {
				"type":        "bool",
				"description": "true if it's proposal for touching blacklist, otherwise false"
			},
			"is_add": {
				"type":        "bool",
				"description": "true if it's proposal for adding, otherwise false"
			},
			"tokens": {
				"type":        "array<string>",
				"description": "describe the tokens to be added / removed from whitelists / blacklists."
			},
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			}
		}
	}`)
}

// RegisterInterfaces register Msg and structs
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpsertTokenRate{},
		&MsgUpsertTokenAlias{},
		&MsgProposalUpsertTokenAlias{},
		&MsgProposalUpsertTokenRates{},
		&MsgProposalTokensWhiteBlackChange{},
	)

	registry.RegisterInterface(
		"kira.gov.Content",
		(*govtypes.Content)(nil),
		&ProposalUpsertTokenAlias{},
		&ProposalUpsertTokenRates{},
		&govtypes.SetPoorNetworkMessagesProposal{},
		&ProposalTokensWhiteBlackChange{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
