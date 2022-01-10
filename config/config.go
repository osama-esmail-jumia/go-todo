package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppMode            string
	AppPort            int
	AppDefaultPageSize int
	AppMaxPageSize     int

	DBDriver   string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func load() Config {
	fang := viper.New()

	fang.SetConfigFile(".env")
	fang.AddConfigPath(".")
	configLocation, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(configLocation)
	}

	fang.AutomaticEnv()
	fang.ReadInConfig() // nolint

	return Config{
		AppMode:            fang.GetString("APP_MODE"),
		AppPort:            fang.GetInt("APP_PORT"),
		AppDefaultPageSize: fang.GetInt("APP_DEFAULT_PAGE_SIZE"),
		AppMaxPageSize:     fang.GetInt("APP_MAX_PAGE_SIZE"),

		DBDriver:   fang.GetString("DB_DRIVER"),
		DBHost:     fang.GetString("DB_HOST"),
		DBPort:     fang.GetInt("DB_PORT"),
		DBUser:     fang.GetString("DB_USER"),
		DBPassword: fang.GetString("DB_PASSWORD"),
		DBName:     fang.GetString("DB_NAME"),
	}
}

var config = load()

func Cfg() *Config { return &config }
