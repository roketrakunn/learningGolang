package main


/**
UNCOMMENT ME DADDy
hello 
these are enums .....
import (
	"fmt"
)
type ServerState int

const ( 

	StateIdle ServerState = iota
	stateConneted
	stateError 
	stateRetrying
)

var stateName = map[ServerState]string{ 
	StateIdle : "idle",
	stateConneted : "Connected",
	stateRetrying : "Retrying",
	stateError : "Error",
}

func (ss ServerState) String() string { 
	return stateName[ss]
}

func transition(s ServerState)  ServerState { 

	switch s{ 

	case StateIdle: 
		return stateConneted

	case stateConneted , stateRetrying: 
		return StateIdle

	case stateError : 
		return  stateError

	default: 
		panic("Unkown state => ")
	}
}

func main(){ 
	
	ns := transition(stateError)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
} */
