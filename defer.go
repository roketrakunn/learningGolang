package main

import (
	"fmt"
)

//Remain the function tp main
func Deferring(){ 
	var sum , n int

	grades := []int{40 , 80 , 78}
	
	for _,val := range  grades { 
		sum +=val
	}
	n = len(grades)	

	defer func () int { 
		avg := sum / n

		fmt.Printf("The avg is %v\n", avg)
		return 1
	}()
}

