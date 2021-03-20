package main

import (
	"LearningGolang/helper"
	"fmt"
)

func main() {
	helper.SayHello("irsan") //Jika nama function nya di awali dengan huruf besar makan function tersebut bisa di akses siapa saja

	//jika di awali huruf kecil tidak bisa di akses siapa saja dan terbatas
	//helper.sayGoodBye("irsan") //error
	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello("irsan"))
}
