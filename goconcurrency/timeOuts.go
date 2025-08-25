package main

// Example program to show how to implement timeouts using select and time.After
// function.
// The program starts two goroutines that simulate work by sleeping for 2 seconds
// before sending a result on a channel. The main function uses select statements
// to wait for either the result from the goroutine or a timeout signal from
// time.After. The first select has a timeout of 1 second, so it will print
// "timeout 1". The second select has a timeout of 3 seconds, so it
// will successfully receive the result from the goroutine and print "result 2".
/**
import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {

    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):

        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2* time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(2 * time.Second):
        fmt.Println("timeout 2")
    }
} */
