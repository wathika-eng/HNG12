package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const baseURL = "https://api.github.com/users"
const PORT = ":8000"

// run the API when called
func main() {
	// load dotenv file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// mux controls traffic
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{username}", getUser)
	mux.HandleFunc("GET /", getUser)
	// delete logs
	defer os.Remove("logs.json")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT)
	// create a server
	log.Printf("Listening on http://localhost%v/\n", PORT)
	go func() {
		http.ListenAndServe(PORT, mux)
	}()
	// when user stops abruptly, ensure logs are also deleted since it's a defer fn
	<-stop
	log.Printf("Shutting down server and deleting logs...\n")
}
