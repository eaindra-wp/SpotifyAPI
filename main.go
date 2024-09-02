package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func getAccessToken() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	tokenUrl := "https://accounts.spotify.com/api/token"

	client := &http.Client{
		Timeout: 9 * time.Second,
	}

	formData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
	}

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the POST request: %v", err)
	}
	defer res.Body.Close()

	var postResponse PostResponse
	if err := json.NewDecoder(res.Body).Decode(&postResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	fmt.Printf("access_token: %v\n", postResponse.AccessToken)
}

func main() {
	getAccessToken()
}
