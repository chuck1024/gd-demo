.PHONY: run docker_run docker

GOPROXY := https://goproxy.cn,direct
GO111MODULE := on

export GO111MODULE
export GOPROXY

default: run

run:
	go run main.go

docker_run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gd-demo main.go
	docker build -t gd-demo .
	rm -rf gd-demo
	docker run -p 10240:10240 -d gd-demo ./server.sh

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gd-demo main.go
	docker build -t registry.cn-chengdu.aliyuncs.com/godog/gd-demo .
	docker tag registry.cn-chengdu.aliyuncs.com/godog/gd-demo registry.cn-chengdu.aliyuncs.com/godog/gd-demo
	docker push registry.cn-chengdu.aliyuncs.com/godog/gd-demo
	rm -rf gd-demo
