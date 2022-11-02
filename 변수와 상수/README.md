# GO lang 변수
## 1. 변수
go 언어에서 변수는 var 키워드를 사용합니다.
var 뒤에 변수명을 작성하고 그 뒤에 type을 적습니다. 아래와 같이 사용합니다.
```golang
var a int
a = 10
```

변수를 선언하면서 값을 할당할 수 있습니다.
```golang
var a float32 = 11 

println(a)  결과 : +1.100000e+001
```

동일 타입의 변수가 여러개 존재할 경우 , 변수를 나열하고 마지막에 type을 지정할 수 있습니다.
또한 복수 변수들이 선언된 상황에서 초기값을 지정할 수 있습니다.
```golang
var a, b, c int
a = 1
b = 2
c = 3

// 복수 변수가 선언된 상황에 초기값 지정
var a,b,c int = 1,2,3
```

## 2. golang 변수 default 값
golang에서 선언된 변수에 초기값을 지정하지 않으면 , Go는 zero value를 지정합니다.
- int : 0
- bool : false
- string : "" ( 빈 문자열 )

## 3. 상수 ( const )
상수는 golang 키워드 const를 사용하여 지정합니다.
사용방법은 변수와 동일한 구조를 가집니다.
```golang
const c int = 10
const s,g,h int = 1,2,3
```

특이하게 , golang에서는 상수를 열거하여 지정할 수 있습니다.
```golang
// const () 사용

const ( // 상수값 열거형 , () 사용
	Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
```

또한 상수값을 열거하였을 경우 , 0부터 순차적으로 iota라는 identifier을 사용할 수 있습니다.
iota가 지정된 상수값에 경우엔 0이 할당되고 , 순차적으로 1씩 증가된 값을 할당받게 됩니다.
```golang


```

## 4. 특이사항
### 4.1 type 추론
변수를 할당할 때 type을 지정해주지 않는다면 , go언어는 type을 예상합니다.
예를들어 아래와 같이 작성해 두었다면 , i에는 정수형이 , s에는 문자열로 type이 지정됩니다.
```golang
var i = 1
var s = "hello"
```

### 4.2 func() 함수 외부에서의 변수 할당
golang은 함수 ( func() ) 외부에서 변수를 할당할 수 있습니다.
```golang
var age int = 26

func main(){
    println(age) // 26
}
```

### 4.3 변수와 상수 허용 범위
golang에서 변수와 상수는 함수 밖에서도 사용할 수 있습니다.