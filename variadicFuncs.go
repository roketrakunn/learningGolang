package main
import "fmt"
func sum(nums...int) int{
	var sum int
	for num := range nums{ 
		sum +=num
	}
	return  sum
}
func main(){ 
	results := sum(1,2,3,4,5)
	fmt.Println(results)
}
