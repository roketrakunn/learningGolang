package main

import "fmt"
type human struct{ 
	height, weight float64
}

func RenameMeTomain() { 
	var m = map[string]human{

		"You":{140, 180}, //cm , kg
		"Me":{180 , 70},
	}
	fmt.Println(m)

	elem, ok := m["Me"]
	fmt.Println(elem)
	fmt.Println(ok)
}
