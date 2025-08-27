package main

import "fmt"
func rangingBaby(){ 

	nums := make(chan string ,2)
	nums <- "One"
	nums <- "Two"
	close(nums)

	for elem := range nums{ 
		fmt.Println("Got ", elem)
	}

}

