package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// all types unexported, but fields must be exported,
// otherwise not visible in json
type status struct {
	API                 string `json:"api"`
	Space               string `json:"space"`
	Logo                string `json:"logo"`
	URL                 string `json:"url"`
	location            `json:"location"`
	contact             `json:"contact"`
	IssueReportChannels []string `json:"issue_report_channels"`
	state               `json:"state"`
}

type location struct {
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type contact struct {
	Phone   string `json:"phone"`
	Irc     string `json:"irc"`
	Twitter string `json:"twitter"`
	Email   string `json:"email"`
	Jabber  string `json:"jabber"`
}

type state struct {
	Open bool `json:"open"`
}

func main() {
	fmt.Println("Webserver running on port 8080")
	http.HandleFunc("/", apiEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiEndpoint(w http.ResponseWriter, r *http.Request) {
	s := status{
		API:   "0.13",
		Space: "Chaos Computer Club Zürich",
		Logo:  "https://stopbuepf.ch/wp-content/uploads/ccczh_logo.png",
		URL:   "http://ccczh.ch",
		location: location{
			Address: "Röschtibachstrasse 26",
			Lat:     47.39306,
			Lon:     8.524826,
		},
		contact: contact{
			Phone:   "+41 79 191 23 70",
			Email:   "presse@ccczh.ch",
			Irc:     "irc://irc.chaostreff.ch/#ccczh",
			Jabber:  "xmpp://ccczh@conference.ccczh.ch",
			Twitter: "@ccczh",
		},
		IssueReportChannels: []string{"Twitter", "E-Mail"},
		state: state{
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
