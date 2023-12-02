package main

import (
	"github.com/kletskovg/accounting/packages/cli/cmd"
	"github.com/kletskovg/accounting/packages/cli/config"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/spf13/viper"
)

func main() {
	config.PrepareViperConfig()
	var mongoUrlConfigValue = viper.Get("MONGODB_URL")
	var mongoUrl, ok = mongoUrlConfigValue.(string)

	if ok {
		db.ConnectDB(mongoUrl)
		cmd.Execute()
	}
}
