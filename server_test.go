package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Return Papper's score", func(t *testing.T) {
		// Assert
		request, _ := http.NewRequest(http.MethodGet, "/players/Paper", nil)
		response := httptest.NewRecorder()
		// Action
		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		// Assert

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		// Assert
		request, _ := http.NewRequest(http.MethodGet, "/palyers/Floyd", nil)
		response := httptest.NewRecorder()

		// Action
		PlayerServer(response, request)
		// Assert
		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
