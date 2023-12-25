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
	DecodeJson()
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

func DecodeJson() {
	myJsonData := []byte(`
	{
		"Price": 299,
		"coursename": "ReactJS Bootcamp",
		"mode": "justcoding.in",
		"tags": ["web-dev","js"]
	}
	`)

	var myCourse course

	checkValid := json.Valid(myJsonData)
	fmt.Println(checkValid)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(myJsonData, &myCourse)
		fmt.Printf(" %#v\n", myCourse)
	} else {
		fmt.Println("That JSON was not valid")
	}

	var myData map[string]interface{}
	json.Unmarshal((myJsonData), &myData)
	fmt.Printf(" %#v\n", myData)

	for key, val := range myData {
		fmt.Printf("Key is %v and valuse is %v and type is %T\n", key, val, val)
	}
}
