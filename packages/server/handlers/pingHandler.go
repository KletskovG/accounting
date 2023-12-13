package handlers

import (
	"io"
	"net/http"
)

type PingHandler struct{}

func (ph PingHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	message := "Service healthy"
	io.WriteString(response, message)
}
