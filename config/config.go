package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoUsername	string `envconfig:"MONGO_USERNAME"`
	MongoPassword string `envconfig:"MONGO_PASSWORD"`
	MongoHost string `envconfig:"MONGO_HOST"`
	MongoPort int `envconfig:"MONGO_PORT"`
}

func FromEnv() (*Config, error) {
	config := Config{}
	if err := config.ReadEnv(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (cfg *Config) ReadEnv() error {
	if err := envconfig.Process("", cfg); err != nil {
		return err
	}

	return nil
}