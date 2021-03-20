package main

import "fmt"

func sumAll(numbers ...int) (int, int) {
	total := 0
	count := 0
	for _, value := range numbers {
		total += value
		fmt.Println("Index ke :", count, "value = ", value)
		count++
	}
	return total, count
}

func main() {
	total, count := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)
	fmt.Println(count)

	/*	slice := []int{1, 2, 3, 4, 5}
		//... untuk slice di masukan jadi parameter
		total = sumAll(slice...)
		fmt.Println(total)*/
}
