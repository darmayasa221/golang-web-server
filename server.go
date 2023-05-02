package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, p.store.GetPlayersScore(player))
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	fmt.Fprint(w, GetPlayersScore(&player))
// }

// func GetPlayersScore(name string) string {
// 	if name == "Paper" {
// 		return "20"
// 	}
// 	if name == "Floyd" {
// 		return "10"
// 	}
// 	return ""
// }
