package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}
func runApp(value int) {
	defer logging()
	fmt.Println("Run apps")
	result := 10 / value
	fmt.Println(result)

}
func main() {
	runApp(10)
}
