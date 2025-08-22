package main

import (
	"fmt"
	"sync"
)

func main(){ 
	var wg sync.WaitGroup
	wg.Add(1) 

	go func(){ 
		defer wg.Done() // signal when finished
		fmt.Println("Hello from a goroutine!")
	}()

	wg.Wait() // wait for goroutine to finish
	fmt.Println("hello from main!")
}

