//interface tipe data abstract , interface berisikan definisi-definisi method

//sebuah kontrak bernama hasname

package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	irsan := Person{
		Name: "Irsan",
	}

	cat := Animal{
		Name: "Millo",
	}
	sayHello(irsan)
	sayHello(cat)
}
