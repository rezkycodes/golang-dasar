package main

import "fmt"

func getHello(name string) string {

	if name == "" {
		return "Hallo bro"
	} else {
		return "Hallo " + name
	}
}
func main() {
	hello := getHello("Rezky")
	fmt.Println(hello)

	fmt.Println(getHello(""))
}
