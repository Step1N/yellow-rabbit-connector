package handlers

import (
	"bytes"
	p "common-server/requestpayload"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//PushRequestHandler handler
func PushRequestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request ")
		if r.Method != "POST" {
			log.Printf("Error in request %v", r.Method)
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var content = &p.PayloadCollection{}
		err := json.NewDecoder(r.Body).Decode(content)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("Error in payload %v", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		bodyBuffer := new(bytes.Buffer)
		json.NewEncoder(bodyBuffer).Encode(content)
		req, err := http.NewRequest("POST", p.YellowRabbitURL+"pubMessage", bodyBuffer)
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
