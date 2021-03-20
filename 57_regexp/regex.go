package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z])([a-z])([a-z])n")
	fmt.Println(regex.MatchString("irsan"))
	fmt.Println(regex.MatchString("irfan"))
	fmt.Println(regex.MatchString("irSan"))

	//find all yg match dengan regex yang kita buat
	search := regex.FindAllString("irsan irfam irfan irzan irwan", -1)
	fmt.Println(search)
}
