package tokens

import (
	"github.com/KiraCore/sekai/x/gov/types"
	"github.com/KiraCore/sekai/x/tokens/keeper"
	types2 "github.com/KiraCore/sekai/x/tokens/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUpsertTokenAliasProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertTokenAliasProposalHandler(keeper keeper.Keeper) *ApplyUpsertTokenAliasProposalHandler {
	return &ApplyUpsertTokenAliasProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUpsertTokenAliasProposalHandler) ProposalType() string {
	return types2.ProposalTypeUpsertTokenAlias
}

func (a ApplyUpsertTokenAliasProposalHandler) Apply(ctx sdk.Context, proposal types.Content) {
	p := proposal.(*types2.ProposalUpsertTokenAlias)

	tokenAlians := types2.NewTokenAlias(p.Symbol, p.Name, p.Icon, p.Decimals, p.Denoms)
	err := a.keeper.UpsertTokenAlias(ctx, *tokenAlians)
	if err != nil {
		panic(err)
	}
}
