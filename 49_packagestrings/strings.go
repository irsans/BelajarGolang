package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Taufik Irsan", "Irsan")) // return value nya bool true or false
	fmt.Println(strings.Split("Taufik Irsan", " "))
	fmt.Println(strings.ToLower("Taufik Irsan"))
	fmt.Println(strings.ToUpper("Taufik Irsan"))
	fmt.Println(strings.ToTitle("Taufik Irsan"))
	fmt.Println(strings.Trim("  Taufik Irsan  ", " "))
	fmt.Println(strings.ReplaceAll("Irsan Irsan Irsan Irsan", "sa", "fa"))
}
