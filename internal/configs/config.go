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
	// Defaults
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.path", "nonProfit-APP.db")
	viper.SetDefault("security.jwt_secret", "CHANGE_ME")
	viper.SetDefault("app.debug", true)

	// Config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Read config (optional)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("⚠️ No config.yaml found, using defaults")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validation
	if cfg.Security.JWTSecret == "CHANGE_ME" {
		fmt.Println("⚠️ WARNING: Using default JWT secret. Change it in config.yaml.")
	}

	return &cfg, nil
}
