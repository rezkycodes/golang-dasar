package main

import "fmt"

func main() {

	//const firstName string = "rezky"
	//const lastName = "Budihartono"
	//const value = 10000

	// OR

	const (
		firstName string = "rezky"
		lastName         = "Budihartono"
		value            = 10000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
