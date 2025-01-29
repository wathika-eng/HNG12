package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ctx    = context.Background()
	client = &http.Client{}
)

// fetch GITHUB AUTH token/API KEY
func getToken() string {
	GITHUB_AUTH_TOKEN, found := os.LookupEnv("GITHUB_AUTH_TOKEN")
	// warn user that no token has been found and that email will be null (incase github's api is rate limiting our IP address)
	if !found || GITHUB_AUTH_TOKEN == "" {
		log.Println(`
		Warning, proceeding without GITHUB AUTH TOKEN, email will default to null
		`)
	}
	return GITHUB_AUTH_TOKEN
}

// w - construct a response for client
// r - info about request (body, header)
func getUser(w http.ResponseWriter, r *http.Request) {
	// ensure data returned is json
	w.Header().Set("Content-Type", "application/json")
	// get data from path params
	userName := r.PathValue("username")
	if userName == "" {
		userName = "wathika-eng"
	}
	GITHUB_AUTH_TOKEN := getToken()
	// format API to include username
	apiUrl := fmt.Sprintf("%v/%v", baseURL, userName)
	fmt.Printf("API URL: %v\n", apiUrl)
	// create a request and utilize the token
	request, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		http.Error(w, "Request not created", http.StatusInternalServerError)
		return
	}
	// set API KEY Token bearer
	request.Header.Set("Authorization", "token "+GITHUB_AUTH_TOKEN)
	request.Header.Set("Accept", "application/vnd.github.v3+json")
	// handle reqeuest
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode == 403 {
		// 403 indicates rate limiting
		http.Error(w, fmt.Sprintf("Error occured while fetching data from GitHub API, status code: %v", resp.StatusCode),
			http.StatusBadRequest)
		return
	}
	// close to save on resources
	defer resp.Body.Close()
	// read data if fetch is successful
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading API data", http.StatusBadRequest)
		return
	}
	logger(respData)
	// indent the response data and filter
	var githubUser map[string]interface{}
	err = json.Unmarshal(respData, &githubUser)
	if err != nil {
		http.Error(w, `{"error": "Failed to parse API response"}`, http.StatusInternalServerError)
		return
	}
	email, emailOk := githubUser["email"].(string)
	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	reposURL, reposOk := githubUser["repos_url"].(string)
	// handle a case where email is private
	if !emailOk || !reposOk {
		// email = "default@email.com"
		email = "EMAIL IS PRIVATE"
		log.Printf("Email is private: %v", email)
	}
	userData := User{
		Email:           email,
		CurrentDatetime: currentTime,
		ReposURL:        reposURL,
	}
	jsonData, err := json.MarshalIndent(userData, "", " ")
	if err != nil {
		http.Error(w, `{"error": "Failed to generate response"}`, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	// return all the data
	// w.Write(respData)
}
