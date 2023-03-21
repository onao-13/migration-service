package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DbName      string
	DbUser      string
	DbPass      string
	DbHost      string
	DbPort      string
	LoadDevData bool
}

var log = logrus.New()

func UploadDevConfig() Config {
	return uploadConfig("dev")
}

func UploadProdConfig() Config {
	return uploadConfig("prod")
}

func uploadConfig(configName string) Config {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName(configName)
	v.SetConfigType("env")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("Error parse config. Error: ", err)
	}

	var config Config

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Error parsing config. Error: ", err)
	}

	return config
}

func (config *Config) GetDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DbUser, config.DbPass, config.DbHost,
		config.DbPort, config.DbName,
	)
}
