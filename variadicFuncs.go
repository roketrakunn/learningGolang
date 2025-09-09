package main

//import "fmt"
//The variadic functions are powerful, everything is powerful in go brother... 
func sum(nums...int) int{
	var sum int
	for num := range nums{ 
		sum +=num
	}
	return  sum
}
/**
UNCOMMENT ME DADDY
func main(){ 
	results := sum(1,2,3,4,5)
	fmt.Println(results)
} */
