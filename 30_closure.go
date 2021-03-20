package main

import "fmt"

func main() {
	//closure adalah ketika kita mengakses variable yang dalam satu scope ketika di ganti maka akan ikut berubah
	counter := 0
	increment := func() {
		counters := counter
		fmt.Println("Increment")
		//counters++
		fmt.Println("ini", counters)
	}

	increment()
	increment()

	fmt.Println(counter)
}
