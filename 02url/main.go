package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.github.io:3000/learn?search=learn&search2=learn2"

func main() {
	//fmt.Println("hii world")
	//fmt.Println("this is a test")
	//fmt.Println(myurl)

	//parsing the url
	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	//another RawQuery and Port()
	//type of the query parameters is url.Values
	fmt.Println("Query:", parsedURL.Query()["search"])

	// adding new query parameters
	query := parsedURL.Query()
	for _, v := range query {
		fmt.Println("Parameter is:", v)
	}
	query.Add("newParam", "newValue")
	parsedURL.RawQuery = query.Encode()

	fmt.Println("Updated URL:", parsedURL.String())

	//create a new url
	newURL := &url.URL{//& sign indicates that the we are passing a reference of url
		Scheme:   "https",
        Host:     "lco.github.io:3000",
        Path:     "/learn",
        RawQuery: "search=learn&search2=learn2",
    }
	fmt.Println("New URL:", newURL.String())
}
