package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/server"
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
