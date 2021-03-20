package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}

	}
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", x)
	}

}
