package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name        string
		Environment string
		Version     string
	}

	Server struct {
		Host string
		Port int
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
		TimeZone string
	}

	Jwt struct {
		Key string
	}
}

var (
	once           sync.Once
	configInstance *Config
)

func LoadConfig(path string) *Config {
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()
		viper.SetEnvPrefix("KYBERMED")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
