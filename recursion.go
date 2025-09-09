package main

//Recursin is still the same brother ... still the same
//import "fmt"
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
/**
UNCOMMENT ME DADDY
func main() {

	fmt.Println(factorial(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-2) + fib(n-1)
	}
	fmt.Println(fib(7))
} */
