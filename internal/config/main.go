package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	DSN string `mapstructure:"dsn"`
}

type Listener struct {
	Addr string `mapstructure:"addr"`
}

type Config struct {
	Database Database `mapstructure:"database"`
	Listener Listener `mapstructure:"listener"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	config := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
