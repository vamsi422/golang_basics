package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso?WSDL="

func main() {
	fmt.Println("web request study.")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
