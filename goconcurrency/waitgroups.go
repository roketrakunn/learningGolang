package main

import (
	"fmt"
	"sync"
	"time"
)

func worrker(id int){

	fmt.Printf("worker %d started." , id)

	fmt.Println("")

	time.Sleep(time.Second)

	fmt.Printf("worker %d finished.", id)
	fmt.Println("")
}


func main(){ 
	var wg sync.WaitGroup
	for i := range 5{ 
		wg.Go(func() {
			worrker(i)
		})
	}
	wg.Wait()
}


