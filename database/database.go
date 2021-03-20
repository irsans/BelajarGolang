package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init di panggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
