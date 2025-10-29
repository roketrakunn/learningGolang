package main

import (
	"fmt"
)

func fibbo(n int ,c chan int){ 
	x, y := 0 , 1
	for range n{ 
		c <- x + y
		x , y = y , x + y
	}
	close(c)
}

func main(){ 
	ch := make(chan int , 10)
	go	fibbo(cap(ch),ch)

	for i:= range ch { 
		fmt.Println(i)
	}
}
