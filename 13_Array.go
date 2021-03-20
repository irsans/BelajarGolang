package main

import "fmt"

func main() {
	var names = [3]string{"taufik", "irsan", "ichan"}
	var umur [2]int
	umur[0] = 15
	umur[1] = 25

	fmt.Println(names[1])
	fmt.Println(umur[1])
	fmt.Println(len(names))
}
