package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	AppPort int `env:"APP_PORT"`

	DBUsername string `env:"MYSQL_USER"`
	DBPassword string `env:"MYSQL_PASSWORD"`
	DBHost     string `env:"MYSQL_HOST"`
	DBPort     int    `env:"MYSQL_PORT"`
	DBName     string `env:"MYSQL_DATABASE"`

	JwtSecret  string `env:"JWT_SECRET"`
	JwtExpired int    `env:"JWT_EXPIRED"`
}

func New() (*Env, error) {
	_ = godotenv.Load()

	config := new(Env)

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
