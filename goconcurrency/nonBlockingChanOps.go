package main


//DAY 3 OF DEDICATING MY LIFE TO GO...

//TO run: Rename nonBlocking() to main()
import (
	"fmt"
)

func nonBlocking(){ 

	message := make(chan string, 1)
	signals :=make ( chan  bool, 1)

	//This is a non blocking recive...if no value is recieved eh default gets executed
	//message <- "There you have it now".. this will be recieved if you uncomment
	select{ 
		case msg := <- message: 
			fmt.Println("Got the message", msg)
		default: 
			fmt.Println("No messages recieved...")
	}

	//This is a non blocking send ... if no value is available to send then we sent 	that there is no value to send lol 
	msg :="Hello..."
	select{ 
		case message<-msg: 
			fmt.Println("Sent message ", msg)
		default: 
			fmt.Println("No message to send...") 
	}

	//Multi way non blocking recieve.. here we attempt to revieve both the message and the signal... if non is recieved the default gets executed....
	select{
		case msg :=<- message: 
			fmt.Println("Recieved message ", msg)
		case sig :=<-signals: 
			fmt.Println("Got the signal ", sig)
		default: 
			fmt.Println("No activity recieved...")	
	}
}
