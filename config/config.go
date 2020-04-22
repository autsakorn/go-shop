package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config defines the properties of environment key
type Config struct {
	MongoUsername string `envconfig:"MONGO_USERNAME"`
	MongoPassword string `envconfig:"MONGO_PASSWORD"`
	MongoHost     string `envconfig:"MONGO_HOST"`
	MongoPort     int    `envconfig:"MONGO_PORT"`
	MongoDatabase string `envconfig:"MONGO_DATABASE"`
}

// FromEnv get all environment keys
func FromEnv() (*Config, error) {
	config := Config{}
	if err := config.readEnv(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (cfg *Config) readEnv() error {
	if err := envconfig.Process("", cfg); err != nil {
		return err
	}

	return nil
}
