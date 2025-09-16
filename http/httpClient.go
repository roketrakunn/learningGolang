package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func renameMeToMain(){ 

	resp , err := http.Get("https://gobyexample.com")
	if err != nil{ 
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response statuts ", resp.Status)

	scanner := bufio.NewScanner(resp.Body) 
	
	for i := 0 ; scanner.Scan() && i <30 ; i++ { 
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err() ; err !=nil{ 
		panic(err)
	}
}

