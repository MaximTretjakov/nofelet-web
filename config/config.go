package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DB           Postgres `env:",prefix=POSTGRES_"`
	Web          WEB      `env:",prefix=WEB_"`
	JWT          JWT      `env:",prefix=JWT_"`
	Debug        bool     `env:"DEBUG"`
	AppNamespace string   `env:"APP_NAMESPACE"`
}

type WEB struct {
	Port              string        `env:"PORT,required"`
	ServerCrt         string        `env:"SERVER_CRT,required"`
	ServerKey         string        `env:"SERVER_KEY,required"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT,default=30s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT,default=30s"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT,default=30s"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT,default=3s"`
}

type JWT struct {
	ValidationKey string `env:"VALIDATION_KEY,required"`
	TTL           int    `env:"TTL,required"`
	Prefix        string `env:"PREFIX,required"`
}

type Postgres struct {
	ConnectionString string `env:"CONNECTION_STRING,required"`
}

func init() {
	_ = godotenv.Load()
}

// NewConfig load current configuration.
func newConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		return nil, err
	}
	return &config, nil
}
