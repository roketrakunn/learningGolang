package main

func fanIn(ch1 , ch2 <-chan string) <- chan string{ 
	out := make(chan string)
	go func(){ 
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




	

	

