package main

import (
	"container/list"
	"strings"
)

func addCarry(a byte, carry1 bool, carry bool) (byte, bool) {
	//fmt.Println("Add Carry:", a, carry)
	if carry {
		if a == '0' {
			return '1', carry1
		} else {
			return '0', true
		}
	} else {
		return a, carry1
	}
}

func addOne(a byte, b byte, carry bool) (byte, bool) {
	//fmt.Println(a, b, carry)
	if a == '0' && b == '0' {
		return addCarry('0', false, carry)
	} else if a == '0' && b == '1' {
		return addCarry('1', false, carry)
	} else if a == '1' && b == '0' {
		return addCarry('1', false, carry)
	} else if a == '1' && b == '1' {
		return addCarry('0', true, carry)
	} else {
		panic("Invalid input")
	}
}

func addBinary(a string, b string) string {
	listA := list.New()
	listB := list.New()
	listR := list.New()
	for i := len(a) - 1; i >= 0; i-- {
		listA.PushBack(a[i])
	}
	for i := len(b) - 1; i >= 0; i-- {
		listB.PushBack(b[i])
	}
	carry := false

	//len_a := listA.Len()

	item_a := listA.Front()
	item_b := listB.Front()
	for {
		if item_a == nil && item_b == nil {
			break
		} else {
			var val1 byte = '0'
			var val2 byte = '0'
			if item_a != nil {
				val1 = item_a.Value.(byte)
				item_a = item_a.Next()
			}
			if item_b != nil {
				val2 = item_b.Value.(byte)
				item_b = item_b.Next()
			}
			val, carry1 := addOne(val1, val2, carry)
			carry = carry1
			//fmt.Println(val, carry)
			listR.PushFront(val)
		}
	}
	if carry {
		var bytec byte = '1'
		listR.PushFront(bytec)
	}

	var build strings.Builder

	for i := listR.Front(); i != nil; i = i.Next() {
		//fmt.Println(i.Value)
		build.WriteByte(i.Value.(byte))
	}

	result := build.String()
	return result
}
