package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/midtrans/midtrans-go"
)

type Config struct {
	ServerPort          string        `envconfig:"PORT" default:"8080"`
	DBHost              string        `envconfig:"DB_HOST" required:"true"`
	DBPort              string        `envconfig:"DB_PORT" required:"true"`
	DBUser              string        `envconfig:"DB_USER" required:"true"`
	DBPassword          string        `envconfig:"DB_PASSWORD" required:"true"`
	DBName              string        `envconfig:"DB_NAME" required:"true"`
	JWTSecretKey        string        `envconfig:"JWT_SECRET_KEY" required:"true"`
	JWTExpiration       time.Duration `envconfig:"JWT_EXPIRATION_HOURS" default:"24h"`
	MidtransServerKey   string        `envconfig:"MIDTRANS_SERVER_KEY" required:"true"`
	MidtransClientKey   string        `envconfig:"MIDTRANS_CLIENT_KEY" required:"true"`
	MidtransEnvironment midtrans.EnvironmentType
	rawMidtransEnv      string `envconfig:"MIDTRANS_ENVIRONMENT" default:"sandbox"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.rawMidtransEnv == "production" {
		cfg.MidtransEnvironment = midtrans.Production
	} else {
		cfg.MidtransEnvironment = midtrans.Sandbox
	}

	return &cfg, nil
}
