package main
import "fmt"
/**
The fibonacci stuff and all

*/

//This is a commet and i am testing ... shit why am i so slow ...


func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		nxt := a
		a, b = b, a+b
		return nxt
	}
}

func renameToMain() {
	fib := fibonacci()
	for i := 0 ; i <= 10 ; i++ {
		fmt.Println(fib())
	}
}
