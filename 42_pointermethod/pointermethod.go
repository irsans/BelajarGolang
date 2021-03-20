package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func main() {
	irsan := Man{"Irsan"}
	irsan.Married()

	fmt.Println(irsan.Name)
}
