package main

import "fmt"

func getCompleteName(firstname string, lastname string, age int) (FirstName string, lastName string, Age int) {
	FirstName = firstname
	lastName = lastname
	Age = age

	return
}
func main() {

	FirstName, lastName, Age := getCompleteName("Taufik", "Irsan", 25)
	fmt.Println(FirstName, lastName, Age)

}
