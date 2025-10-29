package main

import "fmt"
func sum(arr []int , ch chan int){ 
	sum := 0
	for _,val := range arr { 
		sum +=val
	}
	ch <- sum
}

func renameMeToMain() { 


	arr := []int{1,2,3,4}
	c := make(chan int)

	defer close(c) //no need (IN THIS CASE!!)
	go sum(arr[:len(arr)/2],c)
	go sum(arr[len(arr)/2:],c)

	 x , y := <-c, <-c
	fmt.Println(x, y, x+y)
}
