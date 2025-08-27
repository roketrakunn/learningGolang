package main
/** 

uncomment to run
Closing channels

below is an example of closing a channel to indicate that no more values will be sent.

import "fmt"
func main(){ 

	jobs := make(chan int , 5)
	done := make(chan bool)
	go func(){ 
		for {
			j,more := <-jobs

			if more{ 
				fmt.Println("Recieved job ", j)
			}else{ 
				fmt.Println("No more jobs man")
				done <-true
				return
			}
		}
	}()

	for i := range 3{ 
		jobs <- i
		fmt.Println("Sent job ",i)
	}
	close(jobs)
	fmt.Println("All jobs recieved")
	
	<-done	
	_,ok := <-jobs 
	fmt.Println("Successfully delivered all jobs", ok)

} */
