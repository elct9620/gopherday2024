package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const (
	HttpAddrPath = "http.address"
	GrpcAddrPath = "grpc.address"
)

var DefaultSet = wire.NewSet(
	ProvideViper,
	ProvideDefaultConfig,
)

type Config struct {
	HttpAddr string
	GrpcAddr string
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		HttpAddr: v.GetString(HttpAddrPath),
		GrpcAddr: v.GetString(GrpcAddrPath),
	}
}

func ProvideDefaultConfig(v *viper.Viper) *Config {
	v.SetDefault(HttpAddrPath, ":8080")
	v.SetDefault(GrpcAddrPath, ":8081")

	return NewConfig(v)
}

func ProvideViper() (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	return v, nil
}
