package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("this is a test")
	fmt.Println("this is a new line")
}

func PostFormData() {
	// TODO: Implement POST request with form data
	fmt.Println("Posting form data...")
	// Example:
	resp, err := http.PostForm("http://example.com/submit",
		url.Values{
			"name":  {"John"},
			"email": {"john@example.com"},
		})
	if err != nil {
		fmt.Println("Error posting form data:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response headers:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("Response body:")
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(bodyBytes))
	fmt.Println("Form data posted successfully!")
}
