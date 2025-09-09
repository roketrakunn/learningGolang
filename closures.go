package main
//Closures are awesome dude .. everything in go is awesome.
//import "fmt"
func closuring() func () int { 
	dezz := 0
	return func() int {
		dezz +=10
		return dezz
	}
}
/**
func main(){ 

	hol_dezz :=closuring()
	fmt.Println(hol_dezz())
	fmt.Println(hol_dezz())
	fmt.Println(hol_dezz())
	fmt.Println(hol_dezz())

}*/
