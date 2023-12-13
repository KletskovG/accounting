package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {

	t.Run("Healthcheck", func(t *testing.T) {
		mockServer := httptest.NewServer(PingHandler{})
		defer mockServer.Close()

		res, err := http.Get(mockServer.URL)

		if err != nil {
			log.Fatal(err)
		}

		response, err := io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		pingMessage := "Service healthy"

		if string(response) != pingMessage {
			t.Errorf("Ping handler dont return %s message, got %s", pingMessage, string(response))
		}
	})
}
