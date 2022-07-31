package client

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
)

var (
	ModuleBasics = []module.AppModuleBasic{
		auth.AppModuleBasic{},
		authz.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		// TODO: add osmosis governance proposal types here
		// TODO: add other proposal types here
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		feegrant.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ibc.AppModuleBasic{},
	}
)

type ChainClientConfig struct {
	Key            string                  `json:"key" yaml:"key"`
	ChainID        string                  `json:"chain-id" yaml:"chain-id"`
	RPCAddr        string                  `json:"rpc-addr" yaml:"rpc-addr"`
	GRPCAddr       string                  `json:"grpc-addr" yaml:"grpc-addr"`
	AccountPrefix  string                  `json:"account-prefix" yaml:"account-prefix"`
	KeyringBackend string                  `json:"keyring-backend" yaml:"keyring-backend"`
	GasAdjustment  float64                 `json:"gas-adjustment" yaml:"gas-adjustment"`
	GasPrices      string                  `json:"gas-prices" yaml:"gas-prices"`
	KeyDirectory   string                  `json:"key-directory" yaml:"key-directory"`
	Debug          bool                    `json:"debug" yaml:"debug"`
	Timeout        string                  `json:"timeout" yaml:"timeout"`
	BlockTimeout   string                  `json:"block-timeout" yaml:"block-timeout"`
	OutputFormat   string                  `json:"output-format" yaml:"output-format"`
	SignModeStr    string                  `json:"sign-mode" yaml:"sign-mode"`
	Modules        []module.AppModuleBasic `json:"-" yaml:"-"`
}

func (ccc *ChainClientConfig) Validate() error {
	if _, err := time.ParseDuration(ccc.Timeout); err != nil {
		return err
	}
	if ccc.BlockTimeout != "" {
		if _, err := time.ParseDuration(ccc.BlockTimeout); err != nil {
			return err
		}
	}
	return nil
}

func GetCosmosHubConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "notdefault",
		ChainID:        "cosmoshub-4",
		RPCAddr:        "https://cosmoshub-4.technofractal.com:443",
		GRPCAddr:       "https://gprc.cosmoshub-4.technofractal.com:443",
		AccountPrefix:  "cosmos",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "0.01uatom",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetOsmosisConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "notdefault",
		ChainID:        "osmosis-1",
		RPCAddr:        "https://osmosis-1.technofractal.com:443",
		GRPCAddr:       "https://gprc.osmosis-1.technofractal.com:443",
		AccountPrefix:  "osmo",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "0.01uosmo",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetStrideTestnetConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "test",
		ChainID:        "STRIDE",
		RPCAddr:        "http://stride1:26657",
		GRPCAddr:       "http://stride1:9090",
		AccountPrefix:  "stride",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1ustrd",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetGaiaTestnetConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "test",
		ChainID:        "GAIA",
		RPCAddr:        "http://gaia1:26657",
		GRPCAddr:       "http://gaia1:9090",
		AccountPrefix:  "cosmos",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1uatom",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetStrideLocalConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "stridekey",
		ChainID:        "STRIDE",
		RPCAddr:        "http://localhost:26657",
		GRPCAddr:       "http://localhost:9090",
		AccountPrefix:  "stride",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1ustrd",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetGaiaLocalConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "gaiakey",
		ChainID:        "GAIA",
		RPCAddr:        "http://localhost:26557",
		GRPCAddr:       "http://localhost:9080",
		AccountPrefix:  "cosmos",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1uatom",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetJunoLocalConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "junokey",
		ChainID:        "JUNO",
		RPCAddr:        "http://localhost:24657",
		GRPCAddr:       "http://localhost:9030",
		AccountPrefix:  "juno",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1ujuno",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetOsmoLocalConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "osmokey",
		ChainID:        "OSMO",
		RPCAddr:        "http://localhost:23657",
		GRPCAddr:       "http://localhost:9040",
		AccountPrefix:  "osmo",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "1uosmo",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}
