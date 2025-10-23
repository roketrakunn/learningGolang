package main

import "fmt"

//I got owned by arrays so  i am doing them again.
func renameTomain(){ 
	q := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(q)

	r := []bool{true,false,true,false,true,true,false}
	fmt.Println(r)

	s :=[] struct { 
		i int
		b bool
	}{ 
		{2,false},
		{3,false},
		{9,true},
		{8,false},
		{3,true},
		{5,true},
	}
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s))
}

