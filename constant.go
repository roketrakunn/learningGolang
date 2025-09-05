package main

import (
	"fmt"
	"math"
)



func constant(){ 
	//This is a constant

	const string_const = "I am a constant"

	fmt.Println(string_const)
	//This is a constant too
	const n = 50000000

	//Still a constant
	const d = 5e3 / n
	fmt.Println(d)


	//constants have no type.. you can use them the way you wanna
	fmt.Println(int64(n))

	//Or iot is given a type to accordng to the insance it has to be when used
	fmt.Println(math.Sin(n))
}
/** 
UCOMMENT ME
func main(){ 
	constant()
} */

