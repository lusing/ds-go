package main

import (
	"container/list"
	"dsgo/dsgo"
	"fmt"
)

func main() {
	list1 := list.New()
	list1.PushBack(1)
	list1.PushBack(2)
	fmt.Println(list1)
	//fmt.Println("Hello")
	var val = dsgo.AddBinary("11", "1")
	fmt.Println(val)
	dsgo.Test2()
}
