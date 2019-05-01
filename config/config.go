package config

import (
	"github.com/spf13/viper"
)

var Current *Config

type Config struct {
	Addr string
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadConfig(fname string) error {
	viper.SetConfigFile(fname)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Current); err != nil {
		return err
	}

	return nil
}
