package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAddHandler(t *testing.T) {
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

	t.Run("Checks for valid date ", func(t *testing.T) {
		mockServer := httptest.NewServer(AddHandler{})
		defer mockServer.Close()

		// Dont use url.Query().Add() cuz it somehow breaks request
		badDateUrl, _ := url.Parse(mockServer.URL + "?date=2023.12.10&amount=1000")
		t.Logf(badDateUrl.String())
		res, err := http.Get(badDateUrl.String())

		if err != nil {
			t.Errorf(err.Error())
		}

		response, err := io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			t.Errorf(err.Error())
		}

		badDateMessage := "Bad date"

		if string(response) != badDateMessage {
			t.Errorf("Add handler dont return %s message, got %s", badDateMessage, string(response))
		}
	})

	t.Run("Add new transaction", func(t *testing.T) {
		mockServer := httptest.NewServer(AddHandler{})
		defer mockServer.Close()

		// Dont use url.Query().Add() cuz it somehow breaks request
		badDateUrl, _ := url.Parse(mockServer.URL + "?date=2023-12-20&amount=1000&category=other")
		t.Logf(badDateUrl.String())
		res, err := http.Get(badDateUrl.String())

		if err != nil {
			t.Errorf(err.Error())
		}

		responseBody, err := io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			t.Errorf(err.Error())
		}

		if res.StatusCode != http.StatusCreated {
			t.Errorf("Handler didnt response with %d code, got: %d, response: %s", http.StatusCreated, res.StatusCode, string(responseBody))
		}
	})
}
