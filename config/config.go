package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var C Config

type Config struct {
	RPCAddress                string `mapstructure:"rpc_addr"`
	ChainID                   int    `mapstructure:"chain_id"`
	PrivateKey                string `mapstructure:"private_key"`
	MnemonicPhrase            string `mapstructure:"mnemonic_phrase"`
	HDDerivationPath          string `mapstructure:"hd_derivation_path"`
	LiquidStaking             string `mapstructure:"liquid_staking"`
	PledgeOracle              string `mapstructure:"pledge_oracle"`
	TestContract              string `mapstructure:"test_contract"`
	FSKeyStoreDir             string `mapstructure:"dir"`
	StorageProviderRegistry   string `mapstructure:"registry"`
	StorageProviderCollateral string `mapstructure:"collateral"`
	AllocationLimit           int64  `mapstructure:"allocation_limit"`
	MaxPeriod                 int64  `mapstructure:"maxPeriod"`
}

// LoadConfig loads config from files
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("hyperspace")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("collective-sdk")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&C); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the configuration file: %s", err)
	}

	return &C, nil
}
