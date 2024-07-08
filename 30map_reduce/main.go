package main

import (
	"fmt"
	"strings"
)

var words = make(map[string]int)

func main() {

	s := "In the heart of the bustling city, a small garden flourished quietly behind an old brick building, a hidden gem unknown to most. Sunlight dappled through the leaves of ancient oak trees, casting playful shadows on the cobblestone path. The air was fragrant with the scent of jasmine and fresh earth, a natural sanctuary for those who stumbled upon it. Birds chirped melodiously, creating a symphony that contrasted sharply with the distant hum of city traffic. This secret garden, with its tranquil ponds and whispering breezes, offered a peaceful escape from the relentless pace of urban life."
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")

	w := strings.Split(s, " ")
	fmt.Println(len(w))
	map_reduce_func(w)
	fmt.Println(words)
}

func map_reduce_func(s []string) map[string]int {
	for _, v := range s {
		words[v] = words[v] + 1
	}
	return words
}
