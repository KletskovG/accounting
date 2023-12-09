package main

import (
	"net/http"
	"net/url"

	"github.com/kletskovg/accounting/packages/config"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/accounting/server/handlers"
	"github.com/kletskovg/packages/common"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Ping)
	mux.HandleFunc("/ping", handlers.Ping)
	mux.HandleFunc("/list", handlers.ListHander)
	mux.HandleFunc("/add", handlers.AddHandler)
	mux.HandleFunc("/remove", handlers.RemoveHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler)
	mux.HandleFunc("/report", handlers.ReportHandler)

	var port = ":8080"

	logger.Info("Server starting on", port)
	go http.Get(common.TelegramApiUrl + "/done/" + url.PathEscape("Budget server started"))

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
