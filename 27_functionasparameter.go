package main

import "fmt"

//type declaration untuk penamaan func parameter ini seperti filter func (string) string
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

}
func main() {
	sayHelloWithFilter("Irsan", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

}
