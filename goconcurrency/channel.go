package main

import "fmt"

func channel() {
	messages := make(chan string)
	go func() {
		messages <- "I am grooooot"
	}()
	fmt.Println(<-messages)
}
