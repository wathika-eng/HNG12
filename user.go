package main

import (
	"encoding/json"
	"log"
	"os"
)

// user information struct
type User struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	ReposURL        string `json:"repos_url"`
}

// log our API response to a file that will be deleted upon exit
func logger(respData []byte) {
	var JSONlogs map[string]interface{}

	err := json.Unmarshal(respData, &JSONlogs)
	if err != nil {
		log.Printf("error while parsing JSON: %v", err)
		return
	}
	formattedJSON, err := json.MarshalIndent(JSONlogs, "", "  ")
	if err != nil {
		log.Printf("error while formatting JSON: %v", err)
		return
	}
	err = os.WriteFile("logs.json", formattedJSON, 0644)
	if err != nil {
		log.Printf("error while writing log file: %v", err)
	}
}
