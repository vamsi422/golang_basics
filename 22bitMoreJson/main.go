package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json")
	//EncodeJson()
	decodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 499, "LearnCodeOnline.in", "abcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 199, "LearnCodeOnline.in", "abcde123", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "React JS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Input data is correct.")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not VALID")
	}
	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("value value  :%s and type : %T \n", v, v)
		fmt.Printf("key value  :%s and type : %T \n", k, k)

	}

}
