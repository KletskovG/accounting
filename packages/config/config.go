package config

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"github.com/spf13/viper"
)

const (
	ACC_MONGODB_URL        = "ACC_MONGODB_URL"
	ACC_MONGODB_NAME       = "ACC_MONGODB_NAME"
	ACC_MONGODB_COLLECTION = "ACC_MONGODB_COLLECTION"
	ACC_AWS_ACCESS         = "ACC_AWS_ACCESS"
	ACC_AWS_SECRET         = "ACC_AWS_SECRET"
	ACC_AWS_REGION         = "ACC_AWS_REGION"
	ACC_AWS_BUCKET         = "ACC_AWS_BUCKET"
	ACC_MODE               = "ACC_MODE"
)

var requiredVars = []string{
	ACC_MONGODB_COLLECTION,
	ACC_MONGODB_NAME,
	ACC_MONGODB_URL,
}

func init() {
	viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./")
	viper.SetEnvPrefix("ACC")
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
	if _, err := os.Stat(common.CliUserConfigPath); err == nil {
		return
	}

	if len(viper.AllKeys()) > 0 {
		return
	}

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

func ConfigureAwsCLI() error {
	keyIDEnv := GetEnvVariable(ACC_AWS_ACCESS)
	keySecretEnv := GetEnvVariable(ACC_AWS_SECRET)
	regionEnv := GetEnvVariable(ACC_AWS_REGION)

	if keyIDEnv == "" || keySecretEnv == "" || regionEnv == "" {
		return errors.New("Some of AWS ENV variables is missing")
	}

	configureKeyID := exec.Command("aws", "configure", "set", "aws_access_key_id", keyIDEnv)

	if keyIDError := configureKeyID.Run(); keyIDError != nil {
		logger.Info("Cant confugre AWS CLI", keyIDError.Error())
		return keyIDError
	}

	configureAWSSecret := exec.Command("aws", "configure", "set", "aws_secret_access_key", keySecretEnv)

	if AWSSecretError := configureAWSSecret.Run(); AWSSecretError != nil {
		logger.Info("Cant confugre AWS CLI", AWSSecretError.Error())
		return AWSSecretError
	}

	configureRegion := exec.Command("aws", "configure", "set", "region", regionEnv)

	if regionError := configureRegion.Run(); regionError != nil {
		logger.Info("Cant confugre AWS CLI", regionError.Error())
		return regionError
	}

	return nil
}
