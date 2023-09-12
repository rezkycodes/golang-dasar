package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {

		fmt.Println("error dengan message : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp2(error bool) {
	defer endApp()
	if error {
		panic("applikasi error")
	}

	fmt.Println("aplikasi berjalan")
}
func main() {
	runApp2(true)
	fmt.Println("rezky")
}
