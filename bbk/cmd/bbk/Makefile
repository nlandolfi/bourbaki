VERSION = 0.0.9
BUILD_DATE = $(shell date +'%Y-%m-%d_%T')
GIT_SHA = $(shell git rev-parse HEAD)

make:
	go install -ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.GitSHA=$(GIT_SHA)" bbk.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bbk_linux_cmd bbk.go 
