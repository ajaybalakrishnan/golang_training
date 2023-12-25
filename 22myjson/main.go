package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"mode"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Handling")
	EncodeJson()
}

func EncodeJson() {

	mycourses := []course{
		{"ReactJS Bootcamp", 299, "justcoding.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "justcoding.in", "jdfhif333", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 299, "justcoding.in", "abcdjsfhd123", nil},
	}

	// package this data

	// finalJson, err := json.Marshal(mycourses)
	finalJson, err := json.MarshalIndent(mycourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
