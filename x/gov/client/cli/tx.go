package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/KiraCore/sekai/x/gov/types"
	"github.com/KiraCore/sekai/x/staking/client/cli"
)

// define flags
const (
	FlagPermission        = "permission"
	FlagMinTxFee          = "min_tx_fee"
	FlagMaxTxFee          = "max_tx_fee"
	FlagMinValidators     = "min_validators"
	FlagExecName          = "execution_name"
	FlagTxType            = "transaction_type"
	FlagExecutionFee      = "execution_fee"
	FlagFailureFee        = "failure_fee"
	FlagTimeout           = "timeout"
	FlagDefaultParameters = "default_parameters"
	FlagWebsite           = "website"
	FlagMoniker           = "moniker"
	FlagSocial            = "social"
	FlagIdentity          = "identity"
	FlagAddress           = "address"
	FlagWhitelistPerms    = "whitelist"
	FlagBlacklistPerms    = "blacklist"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Custom gov sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewTxCouncilorCmds(),
		NewTxProposalCmds(),
		NewTxRoleCmds(),
		NewTxPermissionCmds(),
		NewTxSetNetworkProperties(),
		NewTxSetExecutionFee(),
	)

	return txCmd
}

// NewTxProposalCmds returns the subcommands of proposal related commands.
func NewTxProposalCmds() *cobra.Command {
	proposalCmd := &cobra.Command{
		Use:                        "proposal",
		Short:                      "Proposal subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	proposalCmd.AddCommand(GetTxProposalAssignPermission())
	proposalCmd.AddCommand(GetTxVoteProposal())
	proposalCmd.AddCommand(GetTxProposalSetNetworkProperty())
	proposalCmd.AddCommand(GetTxProposalSetPoorNetworkMsgs())
	proposalCmd.AddCommand(GetTxProposalCreateRole())
	proposalCmd.AddCommand(GetTxProposalUpsertDataRegistry())

	return proposalCmd
}

// NewTxRoleCmds returns the subcommands of role related commands.
func NewTxRoleCmds() *cobra.Command {
	roleCmd := &cobra.Command{
		Use:                        "role",
		Short:                      "Role subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	roleCmd.AddCommand(GetTxCreateRole())
	roleCmd.AddCommand(GetTxRemoveRole())

	roleCmd.AddCommand(GetTxBlacklistRolePermission())
	roleCmd.AddCommand(GetTxWhitelistRolePermission())
	roleCmd.AddCommand(GetTxRemoveWhitelistRolePermission())
	roleCmd.AddCommand(GetTxRemoveBlacklistRolePermission())

	return roleCmd
}

// NewTxPermissionCmds returns the subcommands of permission related commands.
func NewTxPermissionCmds() *cobra.Command {
	permCmd := &cobra.Command{
		Use:                        "permission",
		Short:                      "Permission subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	permCmd.AddCommand(GetTxSetWhitelistPermissions())
	permCmd.AddCommand(GetTxSetBlacklistPermissions())

	return permCmd
}

func NewTxCouncilorCmds() *cobra.Command {
	councilor := &cobra.Command{
		Use:                        "councilor",
		Short:                      "Councilor subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	councilor.AddCommand(GetTxClaimCouncilorSeatCmd())

	return councilor
}

func GetTxSetWhitelistPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-permission",
		Short: "Whitelists permission into an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			perm, err := cmd.Flags().GetUint32(FlagPermission)
			if err != nil {
				return fmt.Errorf("invalid permissions")
			}

			addr, err := getAddressFromFlag(cmd)
			if err != nil {
				return fmt.Errorf("error getting address: %w", err)
			}

			msg := types.NewMsgWhitelistPermissions(
				clientCtx.FromAddress,
				addr,
				perm,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	setPermissionFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxSetBlacklistPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blacklist-permission",
		Short: "Blacklist permission into an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			perm, err := cmd.Flags().GetUint32(FlagPermission)
			if err != nil {
				return fmt.Errorf("invalid permissions")
			}

			addr, err := getAddressFromFlag(cmd)
			if err != nil {
				return fmt.Errorf("error getting address: %w", err)
			}

			msg := types.NewMsgBlacklistPermissions(
				clientCtx.FromAddress,
				addr,
				perm,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	setPermissionFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewTxSetNetworkProperties is a function to set network properties tx command
func NewTxSetNetworkProperties() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-network-properties",
		Short: "Set network properties",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			minTxFee, err := cmd.Flags().GetUint64(FlagMinTxFee)
			if err != nil {
				return fmt.Errorf("invalid minimum tx fee")
			}
			maxTxFee, err := cmd.Flags().GetUint64(FlagMaxTxFee)
			if err != nil {
				return fmt.Errorf("invalid maximum tx fee")
			}
			minValidators, err := cmd.Flags().GetUint64(FlagMinValidators)
			if err != nil {
				return fmt.Errorf("invalid min validators")
			}

			// TODO: should set more by flags
			msg := types.NewMsgSetNetworkProperties(
				clientCtx.FromAddress,
				&types.NetworkProperties{
					MinTxFee:                    minTxFee,
					MaxTxFee:                    maxTxFee,
					VoteQuorum:                  33,
					ProposalEndTime:             1, // 1min
					ProposalEnactmentTime:       2, // 2min
					EnableForeignFeePayments:    true,
					MischanceRankDecreaseAmount: 10,
					InactiveRankDecreasePercent: 50,      // 50%
					PoorNetworkMaxBankSend:      1000000, // 1M ukex
					MinValidators:               minValidators,
				},
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().Uint64(FlagMinTxFee, 1, "min tx fee")
	cmd.Flags().Uint64(FlagMaxTxFee, 10000, "max tx fee")
	cmd.Flags().Uint64(FlagMinValidators, 2, "min validators")
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxWhitelistRolePermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-permission role permission",
		Short: "Whitelist role permission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			permission, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid permission: %w", err)
			}

			msg := types.NewMsgWhitelistRolePermission(
				clientCtx.FromAddress,
				uint32(role),
				uint32(permission),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewTxSetExecutionFee is a function to set network properties tx command
func NewTxSetExecutionFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-execution-fee",
		Short: "Set execution fee",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			execName, err := cmd.Flags().GetString(FlagExecName)
			if err != nil {
				return fmt.Errorf("invalid execution name")
			}
			txType, err := cmd.Flags().GetString(FlagTxType)
			if err != nil {
				return fmt.Errorf("invalid transaction type")
			}

			execFee, err := cmd.Flags().GetUint64(FlagExecutionFee)
			if err != nil {
				return fmt.Errorf("invalid execution fee")
			}
			failureFee, err := cmd.Flags().GetUint64(FlagFailureFee)
			if err != nil {
				return fmt.Errorf("invalid failure fee")
			}
			timeout, err := cmd.Flags().GetUint64(FlagTimeout)
			if err != nil {
				return fmt.Errorf("invalid timeout")
			}
			defaultParams, err := cmd.Flags().GetUint64(FlagDefaultParameters)
			if err != nil {
				return fmt.Errorf("invalid default parameters")
			}

			msg := types.NewMsgSetExecutionFee(
				execName,
				txType,
				execFee,
				failureFee,
				timeout,
				defaultParams,
				clientCtx.FromAddress,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(FlagExecName, "", "execution name")
	cmd.Flags().String(FlagTxType, "", "execution type")
	cmd.Flags().Uint64(FlagExecutionFee, 10, "execution fee")
	cmd.Flags().Uint64(FlagFailureFee, 1, "failure fee")
	cmd.Flags().Uint64(FlagTimeout, 0, "timeout")
	cmd.Flags().Uint64(FlagDefaultParameters, 0, "default parameters")
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxBlacklistRolePermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blacklist-permission role permission",
		Short: "Blacklist role permissions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			permission, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid permission: %w", err)
			}

			msg := types.NewMsgBlacklistRolePermission(
				clientCtx.FromAddress,
				uint32(role),
				uint32(permission),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxRemoveWhitelistRolePermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-whitelist-permission role permission",
		Short: "Remove whitelist role permissions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			permission, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid permission: %w", err)
			}

			msg := types.NewMsgRemoveWhitelistRolePermission(
				clientCtx.FromAddress,
				uint32(role),
				uint32(permission),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxRemoveBlacklistRolePermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-blacklist-permission role permission",
		Short: "Remove blacklist role permissions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			permission, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid permission: %w", err)
			}

			msg := types.NewMsgRemoveBlacklistRolePermission(
				clientCtx.FromAddress,
				uint32(role),
				uint32(permission),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxCreateRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create role",
		Short: "Create new role",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			msg := types.NewMsgCreateRole(
				clientCtx.FromAddress,
				uint32(role),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxAssignRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assign-role role",
		Short: "Assign new role",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			addr, err := getAddressFromFlag(cmd)
			if err != nil {
				return fmt.Errorf("error getting address: %w", err)
			}

			msg := types.NewMsgAssignRole(
				clientCtx.GetFromAddress(),
				addr,
				uint32(role),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagAddr, "", "the address to set permissions")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(cli.FlagAddr)

	return cmd
}

func GetTxRemoveRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove role",
		Short: "Remove role",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid role: %w", err)
			}

			addr, err := getAddressFromFlag(cmd)
			if err != nil {
				return fmt.Errorf("error getting address: %w", err)
			}

			msg := types.NewMsgRemoveRole(
				clientCtx.FromAddress,
				addr,
				uint32(role),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagAddr, "", "the address to set permissions")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(cli.FlagAddr)

	return cmd
}

// GetTxProposalSetPoorNetworkMsgs defines command to send proposal tx to modify poor network messages
func GetTxProposalSetPoorNetworkMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-poor-network-msgs <messages>",
		Short: "Create a proposal to set a value on a network property.",
		Long: `
		$ %s tx customgov proposal set-poor-network-msgs XXXX,YYY --from=<key_or_address>

		All the message types supported could be added here
			create-role
			assign-role
			remove-role
			...
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			messages := strings.Split(args[0], ",")

			msg := types.NewMsgProposalSetPoorNetworkMessages(
				clientCtx.FromAddress,
				messages,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// GetTxProposalSetNetworkProperty defines command to send proposal tx to modify a network property
func GetTxProposalSetNetworkProperty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-network-property <property> <value>",
		Short: "Create a proposal to set a value on a network property.",
		Long: `
		$ %s tx customgov proposal set-network-property MIN_TX_FEE 100 --from=<key_or_address>

		Available properties:
			MIN_TX_FEE
			MAX_TX_FEE
			VOTE_QUORUM
			PROPOSAL_END_TIME
			PROPOSAL_ENACTMENT_TIME
			ENABLE_FOREIGN_TX_FEE_PAYMENTS
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			property, ok := types.NetworkProperty_value[args[0]]
			if !ok {
				return fmt.Errorf("invalid network property name: %s", args[0])
			}

			value, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid network property value: %w", err)
			}

			msg := types.NewMsgProposalSetNetworkProperty(
				clientCtx.FromAddress,
				types.NetworkProperty(property),
				uint64(value),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxProposalAssignPermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assign-permission permission",
		Short: "Create a proposal to assign a permission to an address.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			perm, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid perm: %w", err)
			}

			addr, err := getAddressFromFlag(cmd)
			if err != nil {
				return fmt.Errorf("error getting address: %w", err)
			}

			msg := types.NewMsgProposalAssignPermission(
				clientCtx.FromAddress,
				addr,
				types.PermValue(perm),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagAddr, "", "the address to set permissions")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(cli.FlagAddr)

	return cmd
}

func GetTxProposalUpsertDataRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upsert-data-registry key hash reference encoding size",
		Short: "Upsert a key in the data registry",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			key := args[0]
			hash := args[1]
			reference := args[2]
			encoding := args[3]
			size, err := strconv.Atoi(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgProposalUpsertDataRegistry(
				clientCtx.FromAddress,
				key,
				hash,
				reference,
				encoding,
				uint64(size),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagAddr, "", "the address to set permissions")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(cli.FlagAddr)

	return cmd
}

func GetTxVoteProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote proposal-id vote-option",
		Short: "Vote a proposal.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid proposal ID: %w", err)
			}

			voteOption, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid vote option: %w", err)
			}

			msg := types.NewMsgVoteProposal(
				uint64(proposalID),
				clientCtx.FromAddress,
				types.VoteOption(voteOption),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// setPermissionFlags sets the flags needed for set blacklist and set whitelist permission
// commands.
func setPermissionFlags(cmd *cobra.Command) {
	cmd.Flags().String(cli.FlagAddr, "", "the address to set permissions")
	cmd.Flags().Uint32(FlagPermission, 0, "the permission")
}

// getAddressFromFlag returns the AccAddress from FlagAddr in Command.
func getAddressFromFlag(cmd *cobra.Command) (sdk.AccAddress, error) {
	addr, err := cmd.Flags().GetString(cli.FlagAddr)
	if err != nil {
		return nil, fmt.Errorf("error getting address")
	}

	bech, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address")
	}

	return bech, nil
}

func GetTxClaimCouncilorSeatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-seat",
		Short: "Claim councilor seat",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moniker, _ := cmd.Flags().GetString(FlagMoniker)
			website, _ := cmd.Flags().GetString(FlagWebsite)
			social, _ := cmd.Flags().GetString(FlagSocial)
			identity, _ := cmd.Flags().GetString(FlagIdentity)
			address, _ := cmd.Flags().GetString(FlagAddress)

			bech32, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimCouncilor(
				moniker,
				website,
				social,
				identity,
				bech32,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(FlagMoniker, "", "the Moniker")
	cmd.Flags().String(FlagWebsite, "", "the Website")
	cmd.Flags().String(FlagSocial, "", "the social")
	cmd.Flags().String(FlagIdentity, "", "the Identity")
	cmd.Flags().String(FlagAddress, "", "the address")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func GetTxProposalCreateRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-role role",
		Short: "Create a proposal to add a new role.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			role, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid perm: %w", err)
			}

			wAsInts, err := cmd.Flags().GetInt32Slice(FlagWhitelistPerms)
			if err != nil {
				return fmt.Errorf("invalid whitelist perms: %w", err)
			}
			whitelistPerms := convertAsPermValues(wAsInts)

			bAsInts, err := cmd.Flags().GetInt32Slice(FlagBlacklistPerms)
			if err != nil {
				return fmt.Errorf("invalid blacklist perms: %w", err)
			}
			blacklistPerms := convertAsPermValues(bAsInts)

			msg := types.NewMsgProposalCreateRole(
				clientCtx.FromAddress,
				types.Role(role),
				whitelistPerms,
				blacklistPerms,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().Int32Slice(FlagWhitelistPerms, []int32{}, "the whitelist value in format 1,2,3")
	cmd.Flags().Int32Slice(FlagBlacklistPerms, []int32{}, "the blacklist values in format 1,2,3")
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// convertAsPermValues convert array of int32 to PermValue array.
func convertAsPermValues(values []int32) []types.PermValue {
	var v []types.PermValue
	for _, perm := range values {
		v = append(v, types.PermValue(perm))
	}

	return v
}
