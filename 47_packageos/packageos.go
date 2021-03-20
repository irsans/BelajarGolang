package main

import (
	"fmt"
	"os"
)

//go run packageos.go taufik irsan
func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	fmt.Println(args[0])
	fmt.Println(args[1])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error,", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	//setelah getenv lakukan set pada cmd jangan terminal visual studio tidak bisa set app_username=(isi dengan yang di inginkan) lakukan di password nya juga
	fmt.Println("ini username:", username)
	fmt.Println("ini pass:", password)
}
