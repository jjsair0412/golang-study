# golang study repository 입니다.

## Go Keyword
golang keyword 목록입니다.
keyword는 변수명 , 상수명 , 함수명 등의 identifier로 사용할 수 없습니다.
```golang
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## kuwon issue
### 1. go: cannot find main module; see ‘go help modules
위 에러 발생하면 , 아래 명령어 에러발생한 golang 경로 터미널로 날린 뒤 작업 시작
```
go env -w GO111MODULE=auto
```