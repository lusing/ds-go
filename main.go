package main

import (
	"container/list"
	"fmt"
)

func main() {
	list1 := list.New()
	list1.PushBack(1)
	list1.PushBack(2)
	fmt.Println(list1)
	//fmt.Println("Hello")
	val := addBinary("11", "1")
	fmt.Println(val)
}
