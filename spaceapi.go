package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	Phone   string `json:"phone"`
	Irc     string `json:"irc"`
	Twitter string `json:"twitter"`
	Email   string `json:"email"`
	Jabber  string `json:"jabber"`
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
		Logo:  "https://stopbuepf.ch/wp-content/uploads/ccczh_logo.png",
		Url:   "http://ccczh.ch",
		Location: Location{
			Address: "Röschtibachstrasse 26",
			Lat:     47.39306,
			Lon:     8.524826,
		},
		Contact: Contact{
			Phone:   "+41 79 191 23 70",
			Email:   "presse@ccczh.ch",
			Irc:     "irc://irc.chaostreff.ch/#ccczh",
			Jabber:  "xmpp://ccczh@conference.ccczh.ch",
			Twitter: "@ccczh",
		},
		Issue_report_channels: []string{"Twitter", "E-Mail"},
		State: State{
			Open: checkDoor(),
		},
	}
	json.NewEncoder(w).Encode(s)
}

func checkDoor() bool {
	t := time.Now()
	hour, _, _ := t.Clock()
	weekday := t.Weekday().String()
	return weekday == "Wednesday" && (hour >= 19 && hour <= 23)
}
