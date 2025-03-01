package cli

import (
	"encoding/json"
	"fmt"

	customgovtypes "github.com/KiraCore/sekai/x/gov/types"
	genutiltypes "github.com/KiraCore/sekai/x/genutil/types"

	customstakingtypes "github.com/KiraCore/sekai/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/KiraCore/sekai/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GenTxClaimCmd(genBalIterator banktypes.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gentx-claim [key_name]",
		Short: "Adds validator into the genesis set",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.JSONMarshaler

			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			_, valPubKey, err := genutil.InitializeNodeValidatorFiles(serverCtx.Config)
			if err != nil {
				return errors.Wrap(err, "failed to initialize node validator files")
			}

			if valPubKeyString, _ := cmd.Flags().GetString(cli.FlagPubKey); valPubKeyString != "" {
				valPubKey, err = types.GetPubKeyFromBech32(types.Bech32PubKeyTypeConsPub, valPubKeyString)
				if err != nil {
					return errors.Wrap(err, "failed to get consensus node public key")
				}
			}

			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(config.GenesisFile())

			name := args[0]
			key, err := clientCtx.Keyring.Key(name)
			if err != nil {
				return errors.Wrapf(err, "failed to fetch '%s' from the keyring", name)
			}

			moniker := config.Moniker
			if m, _ := cmd.Flags().GetString(cli.FlagMoniker); m != "" {
				moniker = m
			}

			website, _ := cmd.Flags().GetString(FlagWebsite)
			identity, _ := cmd.Flags().GetString(FlagIdentity)
			validator, err := customstakingtypes.NewValidator(
				moniker,
				website,
				"social",
				identity,
				types.NewDec(1),
				types.ValAddress(key.GetAddress()),
				valPubKey,
			)
			if err != nil {
				return errors.Wrap(err, "failed to create new validator")
			}

			var stakingGenesisState customstakingtypes.GenesisState
			stakingGenesisState.Validators = append(stakingGenesisState.Validators, validator)
			bzStakingGen := cdc.MustMarshalJSON(&stakingGenesisState)
			appState[customstakingtypes.ModuleName] = bzStakingGen

			var customGovGenState customgovtypes.GenesisState
			cdc.MustUnmarshalJSON(appState[customgovtypes.ModuleName], &customGovGenState)

			// Only first validator is network actor
			networkActor := customgovtypes.NewNetworkActor(
				types.AccAddress(validator.ValKey),
				customgovtypes.Roles{
					uint64(customgovtypes.RoleSudo),
				},
				customgovtypes.Active,
				[]customgovtypes.VoteOption{
					customgovtypes.OptionYes,
					customgovtypes.OptionAbstain,
					customgovtypes.OptionNo,
					customgovtypes.OptionNoWithVeto,
				},
				customgovtypes.NewPermissions(nil, nil),
				1,
			)
			customGovGenState.NetworkActors = append(customGovGenState.NetworkActors, &networkActor)
			appState[customgovtypes.ModuleName] = cdc.MustMarshalJSON(&customGovGenState)

			appGenStateJSON, err := json.Marshal(appState)
			if err != nil {
				return err
			}

			genDoc.AppState = appGenStateJSON

			err = genDoc.ValidateAndComplete()
			if err != nil {
				return err
			}

			err = genDoc.SaveAs(config.GenesisFile())
			if err != nil {
				return err
			}

			fmt.Printf("genesis state updated to include validator\n")

			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	AddValidatorFlags(cmd)

	return cmd
}
