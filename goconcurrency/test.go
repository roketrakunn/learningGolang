package main

import (
	""
	"sync"
)

func fanIn(ch1 , ch2 <-chan string) <- chan string{ 
	out := make(chan string)

	var memMang sync.Mutex //locking 

	
	go func(){ 
		memMang.Lock()
		defer memMang.Unlock()

		for { 
			select{

			case msg :=<- ch1: 
				out <- msg
			case msg :=<- ch2: 
				out <-msg
			}
		}
	}()
	 return out
}




	

	

