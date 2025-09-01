package main

// Tickers are used to perform an action at regular intervals.
// They are similar to timers, but instead of firing once, they fire repeatedly at a specified interval.
// Tickers are useful for tasks that need to be performed periodically, such as polling a resource or updating a UI.
// Tickers can be stopped when they are no longer needed to free up resources. 

//Rename this file to main.go and run it using the command go run main.go 


import (
	"fmt"
	"time"
)

func tickers() {

	ticker := time.NewTicker(500 *time.Millisecond)
	done := make(chan bool)
	
	go func(){
		for{ 
			select{ 
				case <-done: 
					fmt.Println("better luck next time...")
				case t := <-ticker.C: 
					fmt.Println("Ticked at ", t)
			}
		}
	}()

	time.Sleep(1600*time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("kabooom!!!")
}
