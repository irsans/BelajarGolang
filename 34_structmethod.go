package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//penulisan biasa nya seperti ini
func sayHi(customer Customer, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

//penulisan yang di rekomendasikan struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	taufik := Customer{
		Name:    "Irsan",
		Address: "Jakarta",
		Age:     20,
	}

	taufik.sayHello("Taufik")
	sayHi(taufik, "Irsan")
}
