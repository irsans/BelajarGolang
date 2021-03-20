package helper

import "fmt"

var Application = "v10"
var version = 1

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoodBye(name string) {
	fmt.Println("Hellaw", name)
}
