package main

/**
import (
	"fmt"
	"time"
)

// Select statement is used to wait on multiple channel operations.
// It blocks until one of its cases can run, then it executes that case.
// If multiple cases can proceed, one case is chosen at random to execute.
// The default case is executed if no other case is ready.
func main(){
	c1 := make(chan string ,1)
	c2 := make(chan string ,1)


	go func(){ 
		time.Sleep(time.Second*1)
		c1<-"One"
	}()
	go func(){ 
		time.Sleep(2*time.Second)
		c2<-"Two"
	}()
	for range 2{ 
		select{ 
		case msg1:=<-c1: 
			fmt.Println("Hello from ",msg1)	
		case msg2:=<-c2:
			fmt.Println("Hello from ",msg2)		
		default: 
				fmt.Println("No message received")
		}
	}
} */
