package main
//import "fmt"
func zeroval(ival int){ 
	ival = 0
} 
func zeroptr(iptr *int){ 
	*iptr = 0
}
/**UNCOMMENT ME DADDY
func main(){ 
	i :=1 
	fmt.Println("Initail => " , i)

	zeroval(i)
	fmt.Println("Zero val => ",i) //cannot change it because it has no acess to its adress

	zeroptr(&i) //will change it becuase it has is mem adr
	fmt.Println("Zero pointer => ",i)
} */
