package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Taufik") //data paling ujung
	data.PushBack("Irsan")
	data.PushFront("ican")

	//	fmt.Println(data.Front().Next())
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Value)

	//dari depan ke belakang
	for elemet := data.Front(); elemet != nil; elemet = elemet.Next() {
		fmt.Println(elemet.Value)
	}
	//dari belakang ke depan
	for elemet := data.Back(); elemet != nil; elemet = elemet.Prev() {
		fmt.Println(elemet.Value)
	}
}
