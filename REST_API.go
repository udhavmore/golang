/*
Simple REST API to count the site visits and reset the counter
Usage:

	GET /
	POST /reset
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count int

// Simple REST API code to count the Site Visits
func main() {

	server := "localhost:8080"
	fmt.Printf("Starting Server. Listening on: http://%v", server)
	mux := http.NewServeMux()
	count = 0
	mux.HandleFunc("/", index)
	if err := http.ListenAndServe(server, mux); err != nil {
		fmt.Println(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handeGet(w, r)

	case http.MethodPost:
		handlePost(w, r)

	}
}

func handeGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	count++
	msg := "Site Visits: " + strconv.Itoa(count)
	fmt.Fprint(w, msg)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reset" {
		http.NotFound(w, r)
		return
	}
	count = 0
	msg := "Counter was reset"
	fmt.Fprint(w, msg)
}
