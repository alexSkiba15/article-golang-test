package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PostgresDB `yaml:"postgres_db"`
		GRPCServer `yaml:"grpc_server"`
		Server     `yaml:"server"`
	}

	PostgresDB struct {
		Port     string `env-required:"true" yaml:"port"     env:"PG_PORT" env-default:"5432"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		User     string `env-required:"true" yaml:"user"     env:"PG_USER" env-default:"user"`
		DBName   string `env-required:"true" yaml:"db_name"  env:"PG_NAME" env-default:"postgres"`
		Host     string `env-required:"true" yaml:"host"     env:"PG_HOST" env-default:"localhost"`
	}

	Server struct {
		Port string `env-required:"true" yaml:"port" env:"PORT" env-default:"8080"`
	}

	GRPCServer struct {
		Port   string `env-required:"true" yaml:"port" env:"GRPC_PORT" env-default:"8081"`
		Timout string `env-required:"true" yaml:"timout" env:"GRPC_TIMEOUT" env-default:"60s"`
	}
)

func New() *Config {
	cfg := &Config{}
	cwd := projectRoot()
	envFilePath := cwd + ".env"

	err := readEnv(envFilePath, cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func NewPostgresDB() PostgresDB {
	cfg := New()
	return cfg.PostgresDB
}

func readEnv(envFilePath string, cfg *Config) error {
	envFileExists := checkFileExists(envFilePath)

	if envFileExists {
		err := cleanenv.ReadConfig(envFilePath, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	} else {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {

			if _, statErr := os.Stat(envFilePath + ".example"); statErr == nil {
				return fmt.Errorf(
					"missing environmentvariables: %w\n\nprovide all required environment variables or"+
						" rename and update .env.example to .env for convinience", err,
				)
			}

			return err
		}
	}
	return nil
}

func checkFileExists(fileName string) bool {
	envFileExists := false
	if _, err := os.Stat(fileName); err == nil {
		envFileExists = true
	}
	return envFileExists
}

func projectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(b)

	return projectRoot + "/../"
}
