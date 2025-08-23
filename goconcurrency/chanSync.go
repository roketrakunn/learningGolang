package main

import (
	"fmt"
	"time"
)
// This example shows how to use a channel to synchronize a goroutine with the main function.
// The worker function simulates some work by sleeping for a second, then sends a signal on
//Uncomment main to run

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done!")
	done <-true
}
/**
func main(){ 
	done := make(chan bool,1)
	go worker(done)
	<- done
} */ 
