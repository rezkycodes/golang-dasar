package main

import "fmt"

func main() {
	type noKtp string
	type Married bool

	var noKtpR noKtp = "34753498573489"
	var getMarried Married = true

	fmt.Println(noKtpR)
	fmt.Println(getMarried)
}
