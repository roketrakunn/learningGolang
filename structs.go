package main
/**
import (
	"fmt"
)

type person struct{ 
	name string
	age int
}
func newPerson(name string ) *person{ 
	Person :=person{name: name}
	Person.age = 20
	return  &Person
}

func main(){ 

	fmt.Println(person{name: "Me" , age: 20}) //Yeah you can do this too
	fmt.Println(person{name: "Me"}) //Here age is zero ... means i am notb born yet...
	fmt.Println(person{"Me" ,20}) //You can do is .. bad practice tho .. how do you know whether that 20  is my body count or age 
	fmt.Println(&person{name: "Me" , age: 20}) //Yes you can do this as well

	fmt.Println(newPerson("okuhle")) //But this is a more idomatic way to do it .. for like .. i bet you can still find a flow
	s := person{name :"Doom " , age: 500} //Its like you are a god or something...
	fmt.Println(s.name) //This is how you acess fields oy type struct of interest

	sp := &s
	fmt.Println(sp.age) //Its like you hacked yur way in lol .. just kiddingf just basic pointers dude ...outputs  50

	dog := struct{  //This is when you wanna create a unique one time only "object" ... yeah i told you go was awesome
		name string
		age int
	}{ 
		"Brian",
		7,
	}
	fmt.Println(dog)
} */
