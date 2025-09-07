//THE SWITCH STATEMENT

package main

import (
	"fmt"
	"time"
)
func switching(){ 
	
	i := 2
	
	fmt.Print("write " ,i, " as ")
	switch i { 

	case 1: 
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3: 
		fmt.Println("Three")
	}

	switch time.Now().Weekday(){ 

	case time.Sunday , time.Saturday: 
		fmt.Println("No school today.")

	default:
		fmt.Println("Just one more day.")
	}

	whatAmI := func( i any){ 
		switch t := i.(type) { 
			case bool: 
				fmt.Println("Am a boolean")
			case int: 
				fmt.Println("I am an integer")
			default: 
				fmt.Printf("Wow i didnt know about %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(3)
	whatAmI("hello")
}
/**
UNCOMMENT ME
func main(){ 
	switching()
} */
