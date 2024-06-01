package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/linggaaskaedo/go-playground-wire/model/common"
)

func NewAppConfig() common.Configuration {
	var configuration common.Configuration

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Println("Config loaded successfully")

	return configuration
}
