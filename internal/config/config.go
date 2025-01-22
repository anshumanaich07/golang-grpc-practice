package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Protocol string `mapstructure:"protocol"`
	Port     int    `mapstructure:"port"`
}

// read gRPC config
func ReadgRPCConfig() (*GRPCConfig, error) {
	viper.SetConfigName("grpc_configs")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read grpc config")
	}

	conf := &GRPCConfig{}
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal grpc config")
	}

	return conf, nil
}
