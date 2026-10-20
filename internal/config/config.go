package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"database"`

	Security struct {
		JWTSecret string `mapstructure:"jwt_secret"`
	} `mapstructure:"security"`

	App struct {
		Debug bool `mapstructure:"debug"`
	} `mapstructure:"app"`
}

func Load() (*Config, error) {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.path", "nonProfit-APP.db")
	viper.SetDefault("security.jwt_secret", "CHANGE_ME")
	viper.SetDefault("app.debug", true)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	_ = viper.ReadInConfig()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}

	return &cfg, nil
}
