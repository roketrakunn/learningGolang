package main

import "fmt"

func ctrl() {
	//If statement
	if 7%2 == 0 {
		fmt.Println("Even")
	}else { 
		fmt.Println("Odd")
	}

	//More if statements
	if 8 % 4 ==0 { 
		fmt.Println("8 is divisible by 4.")
	}

	//Logical ops are valid too ( == is the reason)
	if 8 % 2 ==0 || 7 % 2 ==0{ 
		fmt.Println("Either 8 or 7 is even.")
	}

	//Statements can precede conditions , how cool is that
	if num := 10 ; num < 0{ 
		fmt.Println(num , " is negative.")
	}else if num < 10{ 
		fmt.Println("Num has mutiple digits.")
	}else { 
		fmt.Println(num, " Has a single digit.")
	}
}
/**
UNCOMMENT ME
func main(){ 
	ctrl()
} */ 
