package config

import (
	"ide/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key" validate:"required,min=32"`
}

type TOTPConfig struct {
	Issuer string `mapstructure:"issuer" validate:"required"`
}

type ServerConfig struct {
	RunEnv types.Env   `mapstructure:"run_env" validate:"required,oneof=dev stg prod"`
	Port   string      `mapstructure:"port" validate:"required,numeric"`
	JWT    *JWTConfig  `mapstructure:"jwt" validate:"required"`
	TOTP   *TOTPConfig `mapstructure:"totp" validate:"required"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required,numeric"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
}

type Config struct {
	Server   *ServerConfig   `mapstructure:"server" validate:"required"`
	Postgres *PostgresConfig `mapstructure:"postgres" validate:"required"`
}

func (c *Config) Validate() error {
	return validator.New().Struct(c)
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error(
			"failed to read config file",
			zap.Error(err),
		)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		zap.L().Error(
			"failed to unmarshal config",
			zap.Error(err),
		)
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		zap.L().Error(
			"config validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info(
		"successfully loaded config",
		zap.String("operation", "config.LoadConfig"),
		zap.Any("run_env", cfg.Server.RunEnv),
		zap.String("port", cfg.Server.Port),
	)

	return &cfg, nil
}
