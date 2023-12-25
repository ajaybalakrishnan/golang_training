package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Welcome to web verb")
	// mygeturl := "http://localhost:8000/get"
	// PerformGerRequest(mygeturl)
	myposturl := "http://localhost:8000/post"
	PerformPostRequest(myposturl)
}

func PerformGerRequest(myurl string) {

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostRequest(myurl string) {
	requestBody := strings.NewReader(`
		{
			"cousrsename":"golang",
			"price": 1000,
			"platform": "online"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
