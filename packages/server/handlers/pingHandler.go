package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Ping(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	message := fmt.Sprint("Service health, %s DB is connected")
	io.WriteString(response, message)
}
