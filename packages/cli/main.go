package main

import (
	"fmt"

	"github.com/kletskovg/accounting/packages/config"
	"github.com/kletskovg/accounting/packages/db"
)

func main() {
	config.PrepareViperConfig()
	var MONGODB_URL, err = config.GetEnvVariable("MONGODB_URL")

	if err != nil {
		panic(err)
	}

	fmt.Println("connecting, ", MONGODB_URL)
	db.ConnectDB(MONGODB_URL)
	// cmd.Execute()
}
