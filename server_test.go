package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{}

	t.Run("Return Papper's score", func(t *testing.T) {
		// Assert
		request := newGetScoreRequest("Paper")
		response := httptest.NewRecorder()
		// Action & Assert
		server.ServerHTTP(response, request)
		assertResponsesBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		// Assert
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// Action & Assert
		server.ServerHTTP(response, request)
		assertResponsesBody(t, response.Body.String(), "10")
	})
}

// hellper function for assert

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponsesBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
