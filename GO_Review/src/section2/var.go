package main

import (
	"fmt"
)

func main() {
	var (
		name   string = "test"
		height int32
		weight float32
	)
	fmt.Println(name, height, weight)

	var (
		food_name string = "떡볶이"
	)
	fmt.Println(food_name)

	const c, d, e = true, 3, "test"

	shortVar1 := 3
	shortVar2 := "test"
	shortVar3 := false

	fmt.Println(shortVar1, shortVar2, shortVar3)

	if i := 10; i < 11 {
		fmt.Println("short var test success")
	}

	const a string = "test1"
	const b int = 32

	const (
		Jan = iota
		Feb
	)

	fmt.Println(Jan, Feb)

	const (
		g = iota * 10
		h
		i
	)

	fmt.Println(g, h, i)

	const (
		_ = iota
		A
		B
		_
		D
	)

	fmt.Println(A, B, D)

}
