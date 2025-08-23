package main

import "fmt"



func main(){ 

	message := make(chan string,2)
	messageUnbufered := make(chan string)

	go func(){ 
		message<-"hello from buffered"
		messageUnbufered<-"hello from unbuffered"
	}()

	fmt.Println(<-message)
	fmt.Println(<-messageUnbufered)

}
