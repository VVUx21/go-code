package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
    fmt.Println("Starting the get request...")
    sendGetRequest("https://lco.github.io:3000/learn?search=learn&search2=learn2")
}

// Function to send GET request

func sendGetRequest(urlString string) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	// Create the GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		fmt.Printf("Error creating GET request: %v\n", err)
		return
	}

	// Send the GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Response body: %s\n", string(bodyBytes))
	fmt.Println("Finished the get request.")
	fmt.Println()
}
