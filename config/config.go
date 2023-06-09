package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DB_USERNAME string `mapstructure:"POSTGRES_DB_USERNAME"`
	POSTGRES_DB_PASSWORD string `mapstructure:"POSTGRES_DB_PASSWORD"`
	POSTGRES_DB_HOST     string `mapstructure:"POSTGRES_DB_HOST"`
	POSTGRES_DB_NAME     string `mapstructure:"POSTGRES_DB_NAME"`
	MIGRATE_PATH         string `mapstructure:"MIGRATE_PATH"`

	SERVER_HOST string `mapstructure:"SERVER_HOST"`

	SALT string `mapstructure:"SALT"`

	ACCESS_TOKEN_EXP  int    `mapstructure:"ACCESS_TOKEN_EXP"`
	REFRESH_TOKEN_EXP int    `mapstructure:"REFRESH_TOKEN_EXP"`
	HS256_SECRET      string `mapstructure:"HS256_SECRET"`

	COINBASE_API string `mapstructure:"COINBASE_API"`
}

func New() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}
	return config, nil
}

func (c *Config) GetPostgresUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", c.POSTGRES_DB_USERNAME, c.POSTGRES_DB_PASSWORD, c.POSTGRES_DB_HOST, c.POSTGRES_DB_NAME)
}
