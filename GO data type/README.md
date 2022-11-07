# GO lang data type
## 1. Go의 data type
GO언어는 다음과 같은 다섯가지 기본 data type을 가지고 있습니다.
1. bool
- true , false인 boolean type
2. 문자열 
- string 
- 한번 생성되면 수정될 수 없습니다.
3. 정수형
- int , int8 , int32 , int64
    - 부호 있는 정수형 타입
    - 각각 8비트 , 32비트 , 64비트 범위를 가집니다.
- uint , uint8 , uint16 , uint64, uintptr
    - 부호 없는 ( unsigned ) 정수형 타입
    - 각각 8비트 , 32비트 , 64비트 범위를 가집니다.
    [관련 문서](https://pyrasis.com/book/GoForTheReallyImpatient/Unit08)
4. Float 및 복소수 타입
- float32 , float64 , complex64 , complex128
- complext는 복소수를 뜻하고 , complex64는 32비트 실수부와 허수부를 나타냅니다. complex128은 64비트 실수부와 허수부를 나타냅니다.
[관련 문서](https://thebook.io/006806/ch03/02/03/)
5. 기타
- byte : uint8과 동일, byte 코드에 사용합니다.
- rune : int32과 동일, 유니코드 코드포인트에 사용합니다.

## 2. 문자열
