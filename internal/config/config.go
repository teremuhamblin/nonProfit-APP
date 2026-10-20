package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabasePath  string
}

func Load() (*Config, error) {
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("database.path", "nonprofit.db")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	_ = viper.ReadInConfig() // ignore si absent

	return &Config{
		ServerAddress: viper.GetString("server.address"),
		DatabasePath:  viper.GetString("database.path"),
	}, nil
}
