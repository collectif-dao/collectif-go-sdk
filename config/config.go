package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

var C Config

type Config struct {
	RPCConfig      map[string]RPCConfig         `mapstructure:"rpc"`
	FSKeyStoreDir  string                       `mapstructure:"dir"`
	DefaultNetwork string                       `mapstructure:"default_network"`
	Addresses      map[string]ContractAddresses `mapstructure:"contracts"`
	Miners         map[string][]string          `mapstructure:"miners"`
	PledgeTimeout  int64                        `mapstructure:"pledge_timeout"`
	LogLevel       string                       `mapstructure:"log_level"`
}

type RPCConfig struct {
	Address  string `mapstructure:"address"`
	ChainID  int    `mapstructure:"chain_id"`
	APIToken string `mapstructure:"api_token"`
}

type ContractAddresses struct {
	LiquidStaking             string `mapstructure:"liquid_staking"`
	StorageProviderRegistry   string `mapstructure:"registry"`
	StorageProviderCollateral string `mapstructure:"collateral"`
}

// LoadConfig loads config from files
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("collective-sdk")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&C); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the configuration file: %s", err)
	}

	if !filepath.IsAbs(C.FSKeyStoreDir) {
		C.FSKeyStoreDir = filepath.Join(filepath.Dir(v.ConfigFileUsed()), C.FSKeyStoreDir)
	}

	return &C, nil
}
