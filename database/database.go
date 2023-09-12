package database

import "fmt"

var connection string

func init() {
	fmt.Println("inits")
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}