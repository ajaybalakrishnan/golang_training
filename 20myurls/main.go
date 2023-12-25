package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?course=golang&paymentid=nsjfiuefn"

func main() {
	fmt.Println("Welcome to handling urls")
	fmt.Println(myurl)

	// parsing the url
	fmt.Println("Parsing url")
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("The type of the result.Query is %T\n", qparam)
	fmt.Println(qparam["course"])
	fmt.Println(qparam["paymentid"])

	for key, val := range qparam {
		fmt.Printf("The key is %v and the value is %v\n", key, val)
	}

	fmt.Println("Constructing URL")
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "learn",
		RawQuery: "course=golang&paymentid=nsjfiuefn",
	}

	anotherUrl := partsOfUrl.String()
	// anotherUrl := string(partsOfUrl)
	fmt.Println(anotherUrl)
}
