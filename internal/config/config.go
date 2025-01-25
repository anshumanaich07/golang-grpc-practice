package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Protocol string `mapstructure:"protocol"`
	Port     int    `mapstructure:"port"`
}

type DBConfig struct {
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  string `mapstructure:"sslmode"`
	TimeZone string `mapstructure:"timezone"`
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

func ReadDBConfig() (*DBConfig, error) {
	viper.SetConfigName("db_config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read db config")
	}

	conf := &DBConfig{}
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal grpc config")
	}

	return conf, nil
}
