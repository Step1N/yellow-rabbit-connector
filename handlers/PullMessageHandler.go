package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	p "common-server/requestpayload"
)

//PullRequestHandler handler
func PullRequestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("PullRequestHandler: Starting parsing request")
		if r.Method != "GET" {
			log.Printf("Error in request %v", r.Method)
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		messageBuffer := new(bytes.Buffer)
		req, err := http.NewRequest("GET", p.YellowRabbitURL+"consumeMessage", messageBuffer)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error in payload %v", err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		err = json.NewEncoder(w).Encode(string(body))
		if err != nil {
			log.Fatalf("Error in payload %v", err)
		}
		w.WriteHeader(http.StatusOK)
	}
}
