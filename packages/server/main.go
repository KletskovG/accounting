package main

import (
	"net/http"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/accounting/server/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/list", handlers.ListHander)
	mux.HandleFunc("/add", handlers.AddHandler)
	mux.HandleFunc("/remove", handlers.RemoveHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler)
	mux.HandleFunc("/report", handlers.ReportHandler)

	var port = ":8080"

	logger.Info("Server starting on", port)
	error := http.ListenAndServe(port, mux)

	if error != nil {
		logger.Error("Error starting server", error)
	}
}
