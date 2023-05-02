package main

import (
	"log"
	"net/http"
)

// main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayersScore(name string) int {
	return 123
}
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
