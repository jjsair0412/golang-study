package main


var a,b = 10,20

var result int

const (
	c = 50
	d = 50
)

func main(){

	arithmetic(c,d)

	relation(c,d)

}

// 산술 연산자 함수
func arithmetic(inputA int, inputB int){

	result = inputA+inputB
	println(result)
	
	result = inputA-inputB
	println(result) 
	
	result = inputA+inputB
	println(result)
	
	result = inputA*inputB
	println(result)
	
	result = inputA/inputB // 나누기 몫
	println(result) // 1
	
	result = inputA%inputB // 나누기 나머지
	println(result) // 0
		
	// 증감연산자
	inputA++
	println(inputA) // 51

	inputA--
	println(inputA) // 50
}

// 관계 연산자 함수
func relation(inputC int , inputD int){
	// true , false로 결과 출력
	println(inputC==inputD)

	println(inputC!=inputD)

	println(inputC>inputD)
}

// 논리연산자 함수
func logical(inputE int , inputF int){
	// AND , OR , NOT 등을 표한함
}