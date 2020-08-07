.PHONY: run docker

GOPROXY := https://goproxy.cn,direct
GO111MODULE := on

export GO111MODULE
export GOPROXY

default: run

run:
	go run main.go

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gd-demo main.go
	docker build -t gd-demo .
	docker run -p 10240:10240 -d gd-demo ./server.sh
