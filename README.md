golang study repository 입니다.

# kuwon issue
## 1. go: cannot find main module; see ‘go help modules
위 에러 발생하면 , 아래 명령어 에러발생한 golang 경로 터미널로 날린 뒤 작업 시작
```
go env -w GO111MODULE=auto
```