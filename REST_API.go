package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Simple REST API code to count the Site Visits
func main() {

	server := "localhost:8080"
	count := 0
	fmt.Printf("Starting Server. Listening on: http://%v", server)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		msg := "Site Visits: " + strconv.Itoa(count)
		fmt.Fprint(w, msg)
	})

	mux.HandleFunc("POST /reset", func(w http.ResponseWriter, r *http.Request) {
		count = 0
		msg := "Counter was reset"
		fmt.Fprint(w, msg)
	})

	if err := http.ListenAndServe(server, mux); err != nil {
		fmt.Println(err.Error())
	}
}
