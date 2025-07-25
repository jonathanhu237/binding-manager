package config

type Config struct {
	ApiServer struct {
		Environment string
		Port        int
	}

	Jwt struct {
		Secret        string
		ExpireMinutes int
	}

	FirstAdmin struct {
		Username string
		Password string
	}

	Postgres struct {
		Username            string
		Password            string
		Host                string
		Port                int
		Db                  string
		PingTimeoutSeconds  int
		QueryTimeoutSeconds int
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	// API Server Configuration
	if cfg.ApiServer.Environment, err = getStringEnv("API_SERVER_ENVIRONMENT"); err != nil {
		return nil, err
	}
	if cfg.ApiServer.Port, err = getIntEnv("API_SERVER_PORT"); err != nil {
		return nil, err
	}

	// JWT Configuration
	if cfg.Jwt.Secret, err = getStringEnv("JWT_SECRET"); err != nil {
		return nil, err
	}
	if cfg.Jwt.ExpireMinutes, err = getIntEnv("JWT_EXPIRE_MINUTES"); err != nil {
		return nil, err
	}

	// First Admin User Configuration
	if cfg.FirstAdmin.Username, err = getStringEnv("FIRST_ADMIN_USERNAME"); err != nil {
		return nil, err
	}
	if cfg.FirstAdmin.Password, err = getStringEnv("FIRST_ADMIN_PASSWORD"); err != nil {
		return nil, err
	}

	// Postgres Configuration
	if cfg.Postgres.Username, err = getStringEnv("POSTGRES_USERNAME"); err != nil {
		return nil, err
	}
	if cfg.Postgres.Password, err = getStringEnv("POSTGRES_PASSWORD"); err != nil {
		return nil, err
	}
	if cfg.Postgres.Host, err = getStringEnv("POSTGRES_HOST"); err != nil {
		return nil, err
	}
	if cfg.Postgres.Port, err = getIntEnv("POSTGRES_PORT"); err != nil {
		return nil, err
	}
	if cfg.Postgres.Db, err = getStringEnv("POSTGRES_DB"); err != nil {
		return nil, err
	}
	if cfg.Postgres.PingTimeoutSeconds, err = getIntEnv("POSTGRES_PING_TIMEOUT_SECONDS"); err != nil {
		return nil, err
	}
	if cfg.Postgres.QueryTimeoutSeconds, err = getIntEnv("POSTGRES_QUERY_TIMEOUT_SECONDS"); err != nil {
		return nil, err
	}

	return &cfg, nil
}
