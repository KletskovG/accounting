package config

import (
	"log"
	"os"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
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
	viper.SetEnvPrefix("MONGODB")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info("Cant read config file, switching to ENV variables")
	}
}

func GetEnvVariable(name string) string {
	CheckUserConfig()
	var value = viper.Get(name)

	if value == nil {
		value = os.Getenv(name)

		if value == "" {
			logger.Info("Cant read ", name, "env variable")
		}
	}

	var valueStr, ok = value.(string)

	if !ok {
		log.Fatal("Variable: ", name, "must be string")
	}

	return valueStr
}

func CheckUserConfig() {
	// fs.Stat(,common.CliUserConfigPath)
	if _, err := os.Stat(common.CliUserConfigPath); err != nil {
		log.Fatal("Cant find user-config file, run 'config command'")
	}

	// TODO: Stopped here, need to parse user config and upload it to viper
}
