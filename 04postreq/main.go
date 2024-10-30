package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl = "https://lco.github.io:3000/learn?search=learn&search2=learn2"

func main() {
	fmt.Println("Starting server for post request")
}

func makepostRequest(url string) {
	requestBody := strings.NewReader(`
		{
	    "search": "learn3",
        "search2": "learn4",
		}
	`)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error while making POST request:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	fmt.Println("Response Body:", string(bodyBytes))
}
