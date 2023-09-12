package main

import "fmt"

func random() interface{} {
	return true
}
func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("integer", value)
	case bool:
		fmt.Println("boolean", value)
	default:
		fmt.Println("uknown", value)
	}
}
