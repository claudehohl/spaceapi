package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Status struct {
	Api                   string `json:"api"`
	Space                 string `json:"space"`
	Logo                  string `json:"logo"`
	Url                   string `json:"url"`
	Location              `json:"location"`
	Contact               `json:"contact"`
	Issue_report_channels []string `json:"issue_report_channels"`
	State                 `json:"state"`
}

type Location struct {
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type Contact struct {
	Twitter string `json:"twitter"`
}

type State struct {
	Open bool `json:"open"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	fmt.Println("Webserver running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	s := Status{
		Api:   "0.13",
		Space: "Chaos Computer Club Zürich",
		Logo:  "http://blafasel/foo.png",
		Url:   "http://ccczh.ch",
		Location: Location{
			Address: "Röschtibachstr",
			Lat:     1.33,
			Lon:     4.77,
		},
		Contact: Contact{
			Twitter: "@ccczh",
		},
		Issue_report_channels: []string{"Twitter", "E-Mail"},
		State: State{
			Open: true,
		},
	}
	json.NewEncoder(w).Encode(s)
}
