package main

import (
	"fmt"
)


func vars(){ 
	//in go you have to  use a variable if you declare it. 

	//Basic string variable 
	var a string
	a = "Hello from a string"
	fmt.Println(a)
	//Int...
	var age int 
	age = 20
	fmt.Println("Yes dude i am ",age)
	//Boolie baby...	
	var single bool
	single = true
	fmt.Println("Dude i am on arch what do you think =>",single)


	//This is however the better way to have em string

	talksTogirls := false
	fmt.Println("Do i speak to girs =>", talksTogirls)
}
/**
UNCOMMET ME !!
func main(){ 
	vars()
} */
