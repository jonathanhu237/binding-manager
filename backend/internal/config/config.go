package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Server struct {
		Environment string
		Port        int
	} `envPrefix:"SERVER_"`

	Jwt struct {
		Secret        string
		ExpireMinutes int
	} `envPrefix:"JWT_"`

	FirstAdmin struct {
		Username string
		Password string
		Email    string
	} `envPrefix:"FIRST_ADMIN_"`

	Postgres struct {
		Username           string
		Password           string
		Host               string
		Port               int
		Db                 string
		PingTimeoutSeconds int
	} `envPrefix:"POSTGRES_"`
}

func LoadConfig() (*Config, error) {
	cfg, err := env.ParseAsWithOptions[Config](env.Options{
		RequiredIfNoDef:       true,
		UseFieldNameByDefault: true,
	})
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
