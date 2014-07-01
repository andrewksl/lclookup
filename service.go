package main

import (
	"encoding/json"
	"github.com/snappymob/lclookup/iso639part1"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	port = ":8088"
)

var (
	p1Map iso639part1.Map
)

func main() {
	// Import file
	f, err := os.Open("data/ISO639-1.tsv")
	if err != nil {
		log.Fatalf(err.Error())
	}
	p1Map, err = iso639part1.GetMap(f)

	// Handlers
	http.HandleFunc("/", home)
	http.HandleFunc("/part1", part1)

	// Listen
	log.Printf("Listening for requests on " + port)
	http.ListenAndServe(port, nil)
}

func part1(w http.ResponseWriter, r *http.Request) {
	c := strings.ToLower(r.FormValue("code"))

	if p1Map[c] != nil {
		e := json.NewEncoder(w)
		e.Encode(p1Map[c])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("html/home.html")
	if err != nil {
		log.Fatalf("Couldn't open home.html: %v", err)
	}
	_, err = io.Copy(w, f)
}
