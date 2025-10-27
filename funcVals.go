package main

import (
	"fmt"
	"math"
)

func compute( fn func(float64, float64) float64) float64{ 
	return fn(3,4)
}

func main(){ 

	hypt := func(x, y float64) float64 { 
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypt(1,2))
	fmt.Println(compute(hypt))

	fmt.Println(compute(math.Pow))

}
