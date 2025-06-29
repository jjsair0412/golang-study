package main

import "fmt"

func main() {

Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("홀수 : ", i)
	}

Loop2:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("i, j : ", i, j)
		}
	}
}
