package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		for {
			// 무한루프
			fmt.Println("infinite loop")
		}
	*/

	loc := []string{"seoul", "busan", "incheon"}

	for index, name := range loc {
		fmt.Println("리스트 위치 : ", index, "값 : ", name)
	}

	for _, name := range loc {
		fmt.Println("리스트 위치 생략, value만 초기화 : ", name)
	}

	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
	}
	fmt.Println(sum2)

	sum3, i2 := 0, 0

	for {
		if i2 > 100 {
			break
		}
		sum3 += i2
		i2++
	}
	fmt.Println(sum3)

	for i4, j := 0, 0; i4 <= 10; i4, j = i4+1, j+10 {
		fmt.Println(i4, j)
	}
}
