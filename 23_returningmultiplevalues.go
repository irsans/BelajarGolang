package main

import "fmt"

//kalau 2 return value harus di kasih tanda kurung
func getFullName() (string, string, int) {
	return "taufik", "irsan", 25
}
func main() {
	//firstName, lastName, Age := getFullName()
	//fmt.Println(firstName, lastName, Age)

	Name, _, _ := getFullName()
	fmt.Println(Name)
}
