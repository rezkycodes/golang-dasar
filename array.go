package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "rezky"
	names[1] = "pradana"
	names[2] = "budihartono"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

}
