package main

import (
	"net/http"

	"github.com/kletskovg/accounting/packages/config"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/accounting/server/handlers"
	"github.com/kletskovg/accounting/server/services"
	"github.com/kletskovg/packages/common"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handlers.PingHandler{})
	mux.Handle("/ping", handlers.PingHandler{})
	mux.Handle("/list", handlers.ListHandler{})
	mux.Handle("/add", handlers.AddHandler{})
	mux.HandleFunc("/remove", handlers.RemoveHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler)
	mux.HandleFunc("/report", handlers.ReportHandler)

	var port = ":8080"

	message := "Budget server starting on " + port
	logger.Info(message)
	go services.Notify(message)

	mode := config.GetEnvVariable(config.ACC_MODE)

	if mode == common.MODE_PROD {
		if configError := config.ConfigureAwsCLI(); configError != nil {
			logger.Error("Cant configure CLI", configError.Error())
			return
		}
	}

	error := http.ListenAndServe(port, mux)

	if error != nil {
		logger.Error("Error starting server", error)
	}
}
