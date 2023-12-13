package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListHandler(t *testing.T) {
	t.Run("Checks for valid amount ", func(t *testing.T) {
		mockServer := httptest.NewServer(AddHandler{})
		defer mockServer.Close()

		// Dont use url.Query().Add() cuz it somehow breaks request
		badAmountUrl, _ := url.Parse(mockServer.URL + "?amount=hello")
		t.Logf(badAmountUrl.String())
		res, err := http.Get(badAmountUrl.String())

		if err != nil {
			t.Errorf(err.Error())
		}

		response, err := io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			t.Errorf(err.Error())
		}

		badAmountMessage := "Bad amount"

		if string(response) != badAmountMessage {
			t.Errorf("Add handler dont return %s message, got %s", badAmountMessage, string(response))
		}
	})
}
