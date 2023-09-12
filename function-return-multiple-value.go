package main

import "fmt"

func getFullName() (string, string) {
	return "rezky", "budihartono"
}
func main() {

	// tanda _ menghiraukan return value
	firstname, _ := getFullName()

	fmt.Println(firstname)
	//fmt.Println(lastname)
}
