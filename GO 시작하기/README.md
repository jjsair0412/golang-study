# GO 시작하기
## 1. golang 설치
[mac 설치 방안](http://golang.site/go/article/202-Mac%EC%97%90-Go-%EC%84%A4%EC%B9%98%ED%95%98%EA%B8%B0)

[linux 설치 방안](http://golang.site/go/article/201-%EB%A6%AC%EB%88%85%EC%8A%A4%EC%97%90-Go-%EC%84%A4%EC%B9%98%ED%95%98%EA%B8%B0)

설치 후 bin 디렉토리를 확인해보면 go.exe파일이 있습니다.
이 파일을 통해 go 프로그램을 컴파일하거나 실행합니다.

go언어는 파일 확장자 .go 를 가집니다.
## 2. run , build
go언어의 명령어는 run과 build가 있습니다.

- go run : go 프로그램을 컴파일과 동시에 실행합니다.
- go build : go 프로그램을 실행파일 .exe 로 생성합니다.

```bash
go run hello.go

go build hello.go
```

## 3. workspace 폴더
Go 프로그래밍을 위해 일반적으로 Workspace 폴더 (작업 폴더)를 생성하는데, 이 폴더 안에는 3개의 서브디렉토리 (bin, pkg, src)를 생성합니다. 

예를 들어, C:\GoApp 디렉토리를 Workspace 폴더로 정했다면, C:\GoApp 안에 bin, pkg, src 서브 폴더를 만들어 줍니다.
```bash
$ pwd
/Users/jjsair0412/Go

$ ls
bin pkg
```

Workspace 폴더를 생성한 후, GOPATH 환경변수에 이 Workspace 폴더 경로를 추가해 줍니다. 
- SET GOPATH=C:\GoApp 처럼 세션별로 할 수 있으나, 주로 시스템 설정에서 시스템 환경변수 혹은 사용자 환경변수로 지정합니다

GOPATH는 하나 이상의 경로를 지정할 수 있습니. 즉, 여러 Workspace가 있는 경우, 이들 경로를 계속 추가할 수 있습니다.

### 3.1 go 환경변수
golang은 중요 환경변수 두 가지를 가집니다.

1. GOROOT
go가 설치된 디렉터리를 가리키며 ( 윈도우 인 경우 C:\GoApp ) Go 실행 파일은 해당 경로의 bin 폴더에,
go 표준 페키지들은 pkg 안에 들어갑니다. ( 시스템 환경변수로 등록됩니다. 관련에러발생시 참고 )

2. GOPATH
go는 표즌 페키지 이외의 third party 페키지나 사용자 정의 페키지들을 GOPATH에서 찾습니다.
복수개의 경로를 지정한 경우 , thrid party 패키지는 처음 경로에 설치됩니다.

golang 환경변수는 go env 명령어로 확인 가능합니다.

```bash
$ go env
...
GOROOT="/usr/local/go"
GOPATH="/Users/jjsair0412/go"
...
```

### 4. etc
http://play.golang.org 

위 웹 사이트에서 golang 프로그램을 편집 , 실행 , 테스트 해 볼 수 있습니다.