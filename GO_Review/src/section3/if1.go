package main

import "fmt"

func main() {
	// 조건문
	// IF문 : 반드시 Boolean 검사 -> 1, 0(자동 형 변환 x)
	// 소괄호 사용 안함

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 dltkd")
	}

	if c := 40; c >= 35 {
		fmt.Println("35 이상")
		c += 20
		fmt.Println(c)
	}

}
