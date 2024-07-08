package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=hdjj38ew99ejww9"

func main() {
	fmt.Println("welcome to handling urls")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparmas := result.Query()
	fmt.Printf("The type of query params are :%T \n", qparmas)
	//fmt.Println(qparmas)
	fmt.Println(qparmas["paymentId"])

	for key, val := range qparmas {
		fmt.Println("key:", key, ", value :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "User=vamshi",
	}
	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)
}
