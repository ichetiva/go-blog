package config

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresDatabase string
}

func NewConfig() *Config {
	return &Config{}
}
