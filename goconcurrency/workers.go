package main

import (
	"fmt"
	"sync"
)

// This example demonstrates the use of goroutines and WaitGroup for concurrency in Go. 
// A goroutine is launched to print a message, and the main function waits for it to complete using a WaitGroup.
// To run this code, simply call the workers() function from your main function.

func workers(){ 
	var wg sync.WaitGroup
	wg.Add(1) 

	go func(){ 
		defer wg.Done() // signal when finished
		fmt.Println("Hello from a goroutine!")
	}()

	wg.Wait() // wait for goroutine to finish
	fmt.Println("hello from main!")
}

