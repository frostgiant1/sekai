package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	customgovtypes "github.com/KiraCore/sekai/x/gov/types"
	"github.com/KiraCore/sekai/x/staking/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	keeper Keeper
}

func NewQuerier(keeper Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

var _ types.QueryServer = Querier{}

func (q Querier) ValidatorByAddress(ctx context.Context, request *types.ValidatorByAddressRequest) (*types.ValidatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	val, err := q.keeper.GetValidator(c, request.ValAddr)
	if err != nil {
		return nil, errors.Wrap(errors.ErrKeyNotFound, err.Error())
	}
	return &types.ValidatorResponse{
		Validator: val,
	}, nil
}

func (q Querier) ValidatorByMoniker(ctx context.Context, request *types.ValidatorByMonikerRequest) (*types.ValidatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)

	val, err := q.keeper.GetValidatorByMoniker(c, request.Moniker)
	if err != nil {
		return nil, errors.Wrap(errors.ErrKeyNotFound, err.Error())
	}

	return &types.ValidatorResponse{
		Validator: val,
	}, nil
}

// Validators implements the Query all validators gRPC method
func (q Querier) Validators(ctx context.Context, request *types.ValidatorsRequest) (*types.ValidatorsResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// validate the provided status, return all the validators if the status is empty
	if request.Status != "" && !strings.EqualFold(request.Status, types.Active.String()) && !strings.EqualFold(request.Status, types.Inactive.String()) && !strings.EqualFold(request.Status, types.Paused.String()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid validator status %s", request.Status)
	}

	store := c.KVStore(q.keeper.storeKey)

	var validators []types.QueryValidator

	validatorStore := prefix.NewStore(store, ValidatorsKey)

	pageRes, err := query.FilteredPaginate(validatorStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var val types.Validator
		err := q.keeper.cdc.UnmarshalBinaryBare(value, &val)
		if err != nil {
			return false, err
		}

		consPubkey, _ := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, val.GetConsPubKey())
		validator := types.QueryValidator{
			Address:    sdk.AccAddress(val.ValKey).String(),
			Valkey:     val.ValKey.String(),
			Pubkey:     consPubkey,
			Proposer:   val.GetConsPubKey().Address().String(),
			Moniker:    val.Moniker,
			Website:    val.Website,
			Social:     val.Social,
			Identity:   val.Identity,
			Commission: val.Commission.String(),
			Status:     val.Status.String(),
			Rank:       val.Rank,
			Streak:     val.Streak,
		}

		if request.Status != "" && !strings.EqualFold(validator.Status, request.Status) {
			return false, nil
		}

		if request.Address != "" && request.Address != validator.Address {
			return false, nil
		}

		if request.Valkey != "" && request.Valkey != validator.Valkey {
			return false, nil
		}

		if request.Pubkey != "" && request.Pubkey != validator.Pubkey {
			return false, nil
		}

		if request.Moniker != "" && request.Moniker != validator.Moniker {
			return false, nil
		}

		if request.Proposer != "" && request.Proposer != validator.Proposer {
			return false, nil
		}

		if request.All || accumulate {
			validators = append(validators, validator)
		}
		return !request.All, nil
	})

	var actors []string
	if request.All {
		for _, actor := range q.keeper.govkeeper.GetNetworkActorsByAbsoluteWhitelistPermission(c, customgovtypes.PermClaimValidator) {
			fmt.Println(actor.Address)
			actors = append(actors, actor.Address.String())
		}
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if request.All {
		pageRes = nil
	}

	response := types.ValidatorsResponse{Validators: validators, Pagination: pageRes, Actors: actors}

	return &response, nil
}
