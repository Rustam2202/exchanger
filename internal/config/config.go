package config

import (
	"flag"
	"log"

	"github.com/Rustam2202/exchanger/internal/logger"
	"github.com/Rustam2202/exchanger/internal/server"
	"github.com/spf13/viper"
)

type Config struct {
	ServerHTTPConfig server.ServerHTTPConfig
	LoggerConfig     logger.LoggerConfig
}

func LoadConfig() *Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
