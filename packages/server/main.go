package main

import (
	"net/http"
	"os"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/accounting/server/handlers"
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
	go http.Get("https://telegram.kletskovg.tech/done/accounting_server_started")

	logger.Info("IS DEV")
	os.Getenv("DEV")

	error := http.ListenAndServe(port, mux)

	if error != nil {
		logger.Error("Error starting server", error)
	}
}
