package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Setup  string `json:"setup"`
	Punch  string `json:"punchline"`
}
//renameMeTomain
func Jsonmain() {
	resp, err := http.Get("https://official-joke-api.appspot.com/jokes/random")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var joke Joke

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		panic(err)
	}
	fmt.Println("Here's a joke for you:")
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punch)
}

