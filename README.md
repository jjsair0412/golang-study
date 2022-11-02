# golang study repository 입니다.
아래 문서를 통해 study한 결과를 정리합니다.
[예제로 배우는 golang](http://golang.site/go/article/1-Go-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D-%EC%96%B8%EC%96%B4-%EC%86%8C%EA%B0%9C)
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