package config

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"github.com/spf13/viper"
)

const (
	MONGODB_URL        = "MONGODB_URL"
	MONGODB_NAME       = "MONGODB_NAME"
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

var requiredVars = []string{MONGODB_COLLECTION, MONGODB_NAME, MONGODB_URL}

func init() {
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./")
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
	if _, err := os.Stat(common.CliUserConfigPath); err != nil {
		logger.Info("Cant find configuration file")
		logger.Info("Please, provide env config variables below (format - <key> <value>): ")
		varsStr := common.ReduceSlice(requiredVars, func(acc, curr string) string {
			acc += " " + curr
			return acc
		}, "")
		logger.Info("Required variables: ", varsStr)
		reader := bufio.NewReader(os.Stdin)
		args, _ := reader.ReadString('\n')
		createUserConfig(args)
	}
}

func createUserConfig(args string) {
	splitResult := strings.Split(args, " ")
	envVars := [][]string{}

	for i := 1; i < len(splitResult); i += 2 {
		envVars = append(envVars, []string{splitResult[i-1], splitResult[i]})
	}

	configFile, err := os.Create(common.CliUserConfigPath)

	if err != nil {
		logger.Error("Cant create config file", err)
	}

	defer configFile.Close()

	for _, envVariable := range envVars {
		envVariableStr := envVariable[0] + "=" + envVariable[1] + "\n"
		configFile.WriteString(envVariableStr)
	}

	logger.Info("Config file created!", configFile.Name())
	viperReInitError := viper.ReadInConfig()
	if viperReInitError != nil {
		logger.Info("Cant read config file, switching to ENV variables")
	}
}
