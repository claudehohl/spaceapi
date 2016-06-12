package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type status struct {
	api      string
	space    string
	logo     string
	url      string
	location struct {
		address string
		lat     float32
		lon     float32
	}
	contact struct {
		twitter string
	}
	issue_report_channels []string
	state                 struct {
		open bool
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	fmt.Println("Webserver running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    //json.NewEncoder(w).Encode(s)
	fmt.Fprintln(w, "Welcome!")
}
