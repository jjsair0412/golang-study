package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := -7

	switch b := 8; {
	case a < 0:
		fmt.Println(a, " 는 음수")
	case a == 0:
		fmt.Println(a, " 는 0과 같음")
	case b > 0:
		fmt.Println(b, " 는 양수")
	}

	// c == ? 같은 값
	switch c := "rust"; c {
	case "go":
		fmt.Println("Go")
	case "c":
		fmt.Println("c")
	case "rust":
		fmt.Println("rust")
	default:
		fmt.Println("일치하는 값 없음")
	}

	switch c := "rust"; c + "lang" {
	case "golang":
		fmt.Println("Golang")
	case "clang":
		fmt.Println("clang")
	case "rustlang":
		fmt.Println("rustlang")
	default:
		fmt.Println("일치하는 값 없음")
	}

	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}
}
