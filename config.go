package main

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	DBConnStr     string `env:"DB_CONN_STR"`
	MigrationsDir string `env:"MIGRATIONS_DIR,required"`
}

func readConfig() (config, error) {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
