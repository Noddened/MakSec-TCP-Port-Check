package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
	BindAddr string `mapstructure:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "localhost:8080",
	}
}

func LoadConfig(logger *slog.Logger) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("configs/")
	viper.AutomaticEnv()

	cfg := NewConfig()
	if err := viper.ReadInConfig(); err != nil {
		logger.Warn("Файл конфига не найден:", "err", err)
		return cfg, nil
	} else {
		logger.Info("Используем конфиг:", "path", viper.ConfigFileUsed())
		if err := viper.Unmarshal(cfg); err != nil {
			return nil, fmt.Errorf("Ошибка сериализации: %w", err)
		}
		return cfg, nil
	}
}
