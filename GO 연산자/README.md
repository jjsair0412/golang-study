# GO lang 연산자
golang에서는 다른 언어와 마찬가지로 아래 종류의 연산자들을 지원 합니다.
1. 산술연산자
2. 관계연산자
3. 논리연산자
4. Bitwise 연산자
5. 할당연산자
6. 포인터연산자

## 1. 산술 연산자
사칙연산 ( + - * / % ) 와 증감 연산자 ( ++ -- ) 를 사용합니다.
```golang
...
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

```

## 2. 관계 연산자
서로의 크기를 비교하거나 동일함을 체크하는데 사용합니다.
ex ) == , != , > , < , >= , <=

```golang
func relation(inputC int , inputD int){
	// true , false로 결과 출력
	println(inputC==inputD)

	println(inputC!=inputD)

	println(inputC>inputD)
}
```

## 3. 논리연산자
AND , OR , NOT 를 표현할 때 사용합니다.
ex ) && , || ,  ! 

## 4. 