package cli_test

import (
	"fmt"
	"testing"

	customgovcli "github.com/KiraCore/sekai/x/gov/client/cli"
	customgovtypes "github.com/KiraCore/sekai/x/gov/types"
	types3 "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/types"

	"github.com/KiraCore/sekai/app"
	"github.com/KiraCore/sekai/simapp"
	"github.com/KiraCore/sekai/testutil/network"
	"github.com/KiraCore/sekai/x/tokens/client/cli"
	tokenstypes "github.com/KiraCore/sekai/x/tokens/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	app.SetConfig()
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	encodingConfig := simapp.MakeEncodingConfig()
	cfg.Codec = encodingConfig.Marshaler
	cfg.TxConfig = encodingConfig.TxConfig

	cfg.NumValidators = 1

	cfg.AppConstructor = func(val network.Validator) servertypes.Application {
		return app.NewInitApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			simapp.MakeEncodingConfig(),
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(types.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestUpsertTokenAliasAndQuery() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	s.WhitelistPermissions(val.Address, customgovtypes.PermUpsertTokenAlias)

	cmd := cli.GetTxUpsertTokenAliasCmd()
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=%s", cli.FlagSymbol, "ETH"),
		fmt.Sprintf("--%s=%s", cli.FlagName, "Ethereum"),
		fmt.Sprintf("--%s=%s", cli.FlagIcon, "myiconurl"),
		fmt.Sprintf("--%s=%d", cli.FlagDecimals, 6),
		fmt.Sprintf("--%s=%s", cli.FlagDenoms, "finney"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	height, err := s.network.LatestHeight()
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(height + 2)
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryTokenAlias()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{"ETH"})
	s.Require().NoError(err)

	var tokenAlias tokenstypes.TokenAlias
	clientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &tokenAlias)

	s.Require().Equal(tokenAlias.Symbol, "ETH")
	s.Require().Equal(tokenAlias.Name, "Ethereum")
	s.Require().Equal(tokenAlias.Icon, "myiconurl")
	s.Require().Equal(tokenAlias.Decimals, uint32(6))
	s.Require().Equal(tokenAlias.Denoms, []string{"finney"})
}

func (s *IntegrationTestSuite) TestUpsertTokenRateAndQuery() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetTxUpsertTokenRateCmd()
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=%s", cli.FlagDenom, "ubtc"),
		fmt.Sprintf("--%s=%f", cli.FlagRate, 0.00001),
		fmt.Sprintf("--%s=%s", cli.FlagFeePayments, "true"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	height, err := s.network.LatestHeight()
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(height + 2)
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryTokenRate()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{"ubtc"})
	s.Require().NoError(err)

	var tokenRate tokenstypes.TokenRate
	clientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &tokenRate)

	s.Require().Equal(tokenRate.Denom, "ubtc")
	s.Require().Equal(tokenRate.Rate, types3.NewDec(10))
	s.Require().Equal(tokenRate.FeePayments, true)
}

func (s *IntegrationTestSuite) TestGetCmdQueryTokenBlackWhites() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryTokenBlackWhites()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{})
	s.Require().NoError(err)

	var blackWhites tokenstypes.TokenBlackWhitesResponse
	clientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &blackWhites)

	s.Require().Equal(blackWhites.Data.Blacklisted, []string{"frozen"})
	s.Require().Equal(blackWhites.Data.Whitelisted, []string{"ukex"})
}

func (s IntegrationTestSuite) TestCreateProposalUpsertTokenRates() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetTxProposalUpsertTokenRatesCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagDenom, "ubtc"),
		fmt.Sprintf("--%s=%f", cli.FlagRate, 0.00001),
		fmt.Sprintf("--%s=%s", cli.FlagFeePayments, "true"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())

	// Vote Proposal
	cmd = customgovcli.GetTxVoteProposal()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("%d", 1), // Proposal ID
		fmt.Sprintf("%d", customgovtypes.OptionYes),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())
}

func (s IntegrationTestSuite) TestCreateProposalUpsertTokenAlias() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetTxProposalUpsertTokenAliasCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagSymbol, "ETH"),
		fmt.Sprintf("--%s=%s", cli.FlagName, "Ethereum"),
		fmt.Sprintf("--%s=%s", cli.FlagIcon, "myiconurl"),
		fmt.Sprintf("--%s=%d", cli.FlagDecimals, 6),
		fmt.Sprintf("--%s=%s", cli.FlagDenoms, "finney"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())

	// Vote Proposal
	out.Reset()
	cmd = customgovcli.GetTxVoteProposal()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("%d", 1), // Proposal ID
		fmt.Sprintf("%d", customgovtypes.OptionYes),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())
}

func (s IntegrationTestSuite) TestTxProposalTokensBlackWhiteChangeCmd() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetTxProposalTokensBlackWhiteChangeCmd()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=true", cli.FlagIsBlacklist),
		fmt.Sprintf("--%s=true", cli.FlagIsAdd),
		fmt.Sprintf("--%s=frozen1", cli.FlagTokens),
		fmt.Sprintf("--%s=frozen2", cli.FlagTokens),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})

	s.Require().NoError(err)
	fmt.Printf("%s", out.String())

	// Vote Proposal
	out.Reset()
	cmd = customgovcli.GetTxVoteProposal()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("%d", 1), // Proposal ID
		fmt.Sprintf("%d", customgovtypes.OptionYes),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, types3.NewCoins(types3.NewCoin(s.cfg.BondDenom, types3.NewInt(100))).String()),
	})
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())
}

// TODO: should add test for GetCmdQueryAllTokenAliases
// TODO: should add test for GetCmdQueryTokenAliasesByDenom
// TODO: should add test for GetCmdQueryAllTokenRates
// TODO: should add test for GetCmdQueryTokenRatesByDenom

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
