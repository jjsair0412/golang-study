package main

var age int = 26

const ( // 상수값 열거형
	Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)

const ( // iota 사용
	apple = iota // 0 
	pizza // 1
 	orange // 2
)



func main()  { // Go는 main 패키지 내의 Entry Point인 main() 함수를 찾아 프로그램을 실행한다.
	var a int // 변수 선언시 var 사용
	a = 10

	var f float32 = 11 // 변수를 선언하면서 값 할당

	var c, d, e int
	c = 1
	d = 2
	e = 3
   
	const s,g,h int = 1,2,3

	println("Hello world")
	println(a) // 10 
	println(f) // +1.100000e+001
	println(c) // 1
	println(d) // 2
	println(e) // 3

	println("상수형",s,g,h) // 상수형 1 2 3

	println("my age is ",age) // my age is  26

	println(Visa,Master,Amex) // Visa MasterCard American Express

	println(apple,pizza,orange) // 0 1 2
}
