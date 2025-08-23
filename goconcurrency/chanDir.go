package main
import "fmt"

// Channel directions
// A channel can be bi-directional, or it can be restricted to send-only or receive-only.
// Here are two functions that demonstrate this concept.
// The pings function is send-only, and the pongs function is receive-only for the ping channel and send-only for the pong channel.// The main function (commented out) shows how to use these functions together.


func pings(ping chan <- string, msg string){ 
	ping <- msg
}

func pongs(ping <- chan string ,pong chan <- string){ 
	msg :=<- ping
	pong <- msg
}
/**
func main(){ 
	ping := make(chan string ,1)
	pong := make(chan string ,1)
	pings(ping, "hello from the ping side")	
	pongs(ping , pong)

	fmt.Println("I the pong was sent by the great pings they are saying ", <-pong)
} */
