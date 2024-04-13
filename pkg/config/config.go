package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	// server
	Port string `env:"PORT" envDefault:"8080"`
	// database
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_DATABASE"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	log.Printf("%+v\n", cfg)
	return cfg, nil
}
