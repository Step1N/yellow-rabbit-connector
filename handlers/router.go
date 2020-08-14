package handlers

import (
	"fmt"
	"net/http"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Context for router
func Context() http.Handler {
	r := mux.NewRouter()

	pushToYellowRabbit := PushRequestHandler()
	pullFromYellowRabbit := PullRequestHandler()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to Messaging Queue")
	})

	r.HandleFunc("/pullParam", pullFromYellowRabbit).Methods("GET")
	r.HandleFunc("/pushParam", pushToYellowRabbit).Methods("POST")

	return gh.CompressHandler(r)
}
