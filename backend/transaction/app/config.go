package app

import "github.com/Ralphbaer/hubla/backend/common"

// Config is the top level configuration struct for entire application
type Config struct {
	EnvName                        string `env:"ENV_NAME"`
	PostgreSQLConnectionString     string `env:"POSTGRES_CONNECTION_STRING"`
	PostgreSQLConnectionStringTest string `env:"POSTGRES_SQL_CONNECTION_TEST"`
	ServerAddress                  string `env:"SERVER_ADDRESS"`
	SpecURL                        string `env:"SPEC_URL"`
}

// NewConfig creates a instance of Config
func NewConfig() *Config {
	cfg := &Config{}
	common.SetConfigFromEnvVars(cfg)
	return cfg
}
