package configfx

import (
	"encoding/json"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Provide(newConfig),
)

type RabbitMQConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Vhost    string `json:"vhost"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Config struct {
	RabbitMQConfig RabbitMQConfig `json:"rabbitmq"`
	DatabaseConfig DatabaseConfig `json:"database"`
}

func newConfig() *Config {
	data, err := os.ReadFile("cmd/config/config.json")
	if err != nil {
		panic(err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
