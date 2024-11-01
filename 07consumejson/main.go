package main

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Timestamp   string   `json:"timestamp"`
	Status      int      `json:"status"`
	Tags        []string `json:"tags"`
}

func main() {
	fmt.Println("hello world")
	fmt.Println("this is a test")
	fmt.Println("this is a new line")
	Decodejson()
}

func Decodejson() {
	jsondatafromweb := []byte(`
		{
            "name":        "Error 1",
            "description": "This is a test error",
            "timestamp":   "2022-01-01T12:00:00Z",
            "status":      500,
            "tags":        ["error", "test"]
        }
	`)
	var myonlinedata map[string]interface{}
	check := json.Valid(jsondatafromweb)
	if !check {
		fmt.Println("Invalid JSON")
	}
	err := json.Unmarshal(jsondatafromweb, &myonlinedata)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		for k, v := range myonlinedata {
			fmt.Printf("%v: %v and type is : %T\n", k, v, v)
		}

		fmt.Printf("Decoded JSON: %+v\n", myonlinedata)
	}
	fmt.Println("end of program")
	fmt.Println()
}
