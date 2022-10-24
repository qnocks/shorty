package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type (
	Config struct {
		Server ServerConfig
		Redis  RedisConfig
	}

	ServerConfig struct {
		Port string `env:"PORT" env-default:"8080"`
	}

	RedisConfig struct {
		Host     string `env:"REDIS_HOST" env-default:"localhost"`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
		Password string `env:"REDIS_PASS" env-default:""`
	}
)

var instance *Config

func GetConfig() *Config {
	once := sync.Once{}
	once.Do(func() {
		instance = &Config{}

		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("error during loading enviroment variables: %s\n", err.Error())
		}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalf("error during mapping enviroment variables: %s\n", help)
		}
	})

	return instance
}
