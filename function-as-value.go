package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}
func main() {
	sayGoodBye := getGoodBye

	fmt.Println(sayGoodBye("Rezky"))
}
