## go-book-library
(private project) 도서 관리 API 시스템

## 실행 환경 구성

### 1. 의존 모듈 설치(터미런에서 실행)

#### 1.1. echo 설치

$ go get -u github.com/labstack/echo/...

#### 1.2. mysql 드라이버 설치

$ go get github.com/go-sql-driver/mysql

#### 1.3. zap 로거 설치
$ go get -u go.uber.org/zap


### 2. 환경 변수 셋팅

윈도우 환경 변수 GOPATH를 추가하고, 프로젝트 루트 디렉토리를 등록
