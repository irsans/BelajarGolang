package main

import "fmt"

func main() {

	type NoKTP string

	var KTP NoKTP = "12345678"
	fmt.Println(KTP)
	fmt.Println(NoKTP("44444444"))
}
