package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// PaymentResponse - The payment response
type PaymentResponse struct {
	Result bool `json:"result"`
}

// With inspiration from https://github.com/kelseyhightower/app-healthz
//  Not that this service has dependencies, yet...
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	x := "{\"status\": \"Ready and waiting\", \"up\": true}"
	fmt.Fprint(w, x)
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not supported", 400)
		return
	}

	// Consider if we should de-serialize request
	//  This would only be useful if we wanted to be consistent with our response

	// Consider if we should seed the random generator....
	result := (rand.Intn(2) == 0)
	response := &PaymentResponse{Result: result}

	x, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Woops", 501)
	}

	fmt.Fprint(w, string(x))
}

func main() {
	log.Println("Starting payment app")

	http.HandleFunc("/", payHandler)
	http.HandleFunc("/healthz", healthzHandler)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
