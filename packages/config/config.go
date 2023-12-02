package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	MONGODB_URL        = "MONGODB_URL"
	MONGODB_NAME       = "MONGODB_NAME"
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

func init() {
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetEnvVariable(name string) string {
	var value = viper.Get(name)

	if value == nil {
		log.Fatal("Cant read ", name, "config variable")
	}

	var valueStr, ok = value.(string)

	if !ok {
		log.Fatal("Variable: ", name, "must be string")
	}

	return valueStr
}
