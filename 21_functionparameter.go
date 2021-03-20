package main

import "fmt"

func sayeHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
func main() {
	firstname := "Taufik"
	lastName := "Irsan"
	sayeHelloTo(firstname, lastName)
}
