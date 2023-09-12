package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mt. " + man.Name
}
func main() {
	rezky := Man{"Rezky"}
	rezky.Married()

	fmt.Println(rezky.Name)
}
