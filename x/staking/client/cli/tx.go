package cli

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/KiraCore/sekai/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	customstakingtypes "github.com/KiraCore/sekai/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagWebsite   = "website"
	FlagSocial    = "social"
	FlagIdentity  = "identity"
	FlagComission = "commission"
	FlagValKey    = "validator-key"
)

func GetTxClaimValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-validator-seat",
		Short: "Claim validator seat to become a Validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			serverCtx := server.GetServerContextFromCmd(cmd)

			moniker, _ := cmd.Flags().GetString(FlagMoniker)
			website, _ := cmd.Flags().GetString(FlagWebsite)
			social, _ := cmd.Flags().GetString(FlagSocial)
			identity, _ := cmd.Flags().GetString(FlagIdentity)
			comission, _ := cmd.Flags().GetString(FlagComission)

			var (
				valPubKey cryptotypes.PubKey
			)
			if valPubKeyString, _ := cmd.Flags().GetString(cli.FlagPubKey); valPubKeyString != "" {
				valPubKey, err = types.GetPubKeyFromBech32(types.Bech32PubKeyTypeConsPub, valPubKeyString)
				if err != nil {
					return errors.Wrap(err, "failed to get consensus node public key")
				}
			} else {
				_, valPubKey, err = genutil.InitializeNodeValidatorFiles(serverCtx.Config)
				if err != nil {
					return errors.Wrap(err, "failed to initialize node validator files")
				}
			}

			comm, err := types.NewDecFromStr(comission)
			val := types.ValAddress(clientCtx.GetFromAddress())

			msg, err := customstakingtypes.NewMsgClaimValidator(moniker, website, social, identity, comm, val, valPubKey)
			if err != nil {
				return fmt.Errorf("error creating tx: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	AddValidatorFlags(cmd)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// GetTxProposalUnjailValidatorCmd implement cli command for MsgUpsertTokenAlias
func GetTxProposalUnjailValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-unjail-validator hash reference",
		Short: "Creates an proposal to unjail validator (the from address is the validator)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			hash := args[0]
			reference := args[1]

			msg := customstakingtypes.NewMsgProposalUnjailValidator(
				clientCtx.FromAddress,
				hash,
				reference,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
