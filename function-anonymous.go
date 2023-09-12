package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

//	func blacklistAdmin(name string) bool {
//		return name == "admin"
//	}
func main() {
	//registerUser("admin", blacklistAdmin)

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("rezky", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
