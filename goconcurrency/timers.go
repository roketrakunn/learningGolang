package main

import (
	"fmt"
	"time"
)

/**
 * Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
 * Timers can be used for timeouts and to delay execution.
*/



func defuseTheBomb(){ 

	timer1 := time.NewTimer(time.Second*2)

	<-timer1.C 
	fmt.Println("Timer 1 fired!!")
	
	timer2 := time.NewTimer(5*time.Second)

	fmt.Println("Bomb planted")
	go func(){
		<-timer2.C
		fmt.Println("BOOM!!")
	}()
	// wait for 3 seconds before defusing the bomb
	time.Sleep(time.Second*3)

	fmt.Println("Defusing the bomb...")
	// defuse the bomb
	stop2 := timer2.Stop()

	time.Sleep(time.Second*3)

	if stop2{ fmt.Println("Successfully defused the bomb")}
	time.Sleep(time.Second*1)
	fmt.Println("Congratulations you have saved the world!")
	
}
