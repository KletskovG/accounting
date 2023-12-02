package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func PrepareViperConfig() {
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// TODO: make name paramters some sort of "set of strings"
func GetEnvVariable(name string) (string, error) {
	var value = viper.Get(name)

	var valueStr, ok = value.(string)
	var errorResult error = nil

	if !ok {
		errorResult = errors.New("Variable " + name + "is not a string")
	}

	return valueStr, errorResult
}
