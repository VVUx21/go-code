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
	// fmt.Println("hello world")
	// fmt.Println("this is a test")
	// fmt.Println("this is a new line")

	Createjsondata()
}

func Createjsondata() {
	// error := Error{
	// 	Name:        "Error 1",
	// 	Description: "This is a test error",
	// 	Timestamp:   "2022-01-01T12:00:00Z",
	// 	Status:      500,
	// 	Tags:        []string{"error", "test"},
	// }

	error := []Error{
		{
            Name:        "Error 1",
            Description: "This is a test error",
            Timestamp:   "2022-01-01T12:00:00Z",
            Status:      500,
            Tags:        []string{"error", "test"},
        },
        {
            Name:        "Error 2",
            Description: "Another test error",
            Timestamp:   "2022-01-02T13:00:00Z",
            Status:      404,
            Tags:        nil,
        },
    }
	// jsonData, err := json.Marshal(error)
	jsonData, err := json.MarshalIndent(error, "", "\t")
    if err!= nil {
        fmt.Println("Error marshaling JSON:", err)
    }

    fmt.Println(string(jsonData))
}
