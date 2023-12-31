package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PostgresConfig PostgresDB
	}

	PostgresDB struct {
		Port     string `env-required:"true" env:"PG_PORT" env-default:"5432"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		User     string `env-required:"true" env:"PG_USER" env-default:"user"`
		DBName   string `env-required:"true" env:"PG_NAME" env-default:"postgres"`
		Host     string `env-required:"true" env:"PG_HOST" env-default:"localhost"`
	}

	Server struct {
		Port string `env-required:"true" env:"PORT" env-default:"8080"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
