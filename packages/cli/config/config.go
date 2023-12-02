package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func PrepareViperConfig() {
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
