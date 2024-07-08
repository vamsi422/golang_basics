package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb video.")
	//performGetRequest("http://localhost:8000/get")
	//performPostRequest("http://localhost:8000/post")
	performPostFormRequest("http://localhost:8000/postform")
}

func performGetRequest(myUrl string) {

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("result is:", response.StatusCode)
	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println("content in string format:", responseString.String())
	fmt.Println("response is:", responseString)
	fmt.Println("CONTENT IS :", content)

	// fmt.Println("content length is:", response.ContentLength)
	// fmt.Println("content inside body is :", string(content))

	defer response.Body.Close()
}

func performPostJsonRequest(myUrl string) {

	requestBody := strings.NewReader(`{
		"course":"react js",
		"price":"12",
		"current discount":"10%",
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("response is :", string(content))
}

func performPostFormRequest(myUrl string) {
	data := url.Values{}
	data.Add("firstName", "Vamshi")
	data.Add("lastName", "Krishna")
	data.Add("Email", "Vamshi@dev")
	data.Add("status", "inActive")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(string(content))

}
