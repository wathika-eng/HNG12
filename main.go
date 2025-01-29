package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT = ":8000"

// run the API when called
func main() {
	// mux controls traffic
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getUser)

	// create a server
	log.Printf("Listening on http://localhost%v/\n", PORT)
	http.ListenAndServe(PORT, mux)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// ensure data is in json format
	w.Header().Set("Content-Type", "application/json")
	userData := map[string]interface{}{
		"email":            "wathika02@gmail.com",
		"current_datetime": time.Now().UTC().Format(time.RFC3339),
		"github_url":       "https://github.com/wathika-eng/simple_github_user_api",
	}

	data, err := json.MarshalIndent(userData, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	data = []byte(fmt.Sprintf("%s\n", data))
	// 200, okay
	w.WriteHeader(http.StatusOK)
	// json data
	w.Write(data)
}
