# gd-demo

## 目录结构说明

|目录/文件名称   | 说明 | 描述
|---|---|---
|`app`           | 业务逻辑层 | 所有的业务逻辑存放目录。
| - `api`        | 业务接口   | 接收/解析用户输入参数的入口/接口层。
| - `model`      | 数据模型   | 数据管理层，仅用于操作管理数据，如数据库操作。
| - `service`    | 逻辑封装   | 业务逻辑封装层，实现特定的业务需求，可供不同的包调用。
|`boot`          | 初始化包   | 用于项目初始化参数设置。
|`conf`          | 配置管理   | 所有的配置文件存放目录。
|`docker`        | 镜像文件   | Docker镜像相关依赖文件，脚本文件等等。
|`docs`          | swag文件  | swagger文件。
|`document`      | 项目文档   | Document项目文档，如: 设计文档、帮助文档等等。
|`library`       | 公共库包   | 公共的功能封装包，往往不包含业务需求实现。
|`router`        | 路由注册   | 用于路由统一的注册管理。
|`Dockerfile`    | 镜像描述   | 云原生时代用于编译生成Docker镜像的描述文件。
|`go.mod`        | 依赖管理   | 使用`Go Module`包管理的依赖描述文件。
|`main.go`       | 入口文件   | 程序入口文件。
|`Makefile`      | make文件  | makefile。

## swag

生成 swagger文件

```linux
swag init -g=./main.go
```

访问 swagger 

http://127.0.0.1:10240/seagger/index.html

## docker 运行

```bash
# 编译程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gd-demo main.go

# docker build 镜像
docker build -t gd-demo .

# docker 运行程序
docker run -p 10240:10240 -d gd-demo ./server.sh

# docker ps
docker ps

# kill 该进程 xxxxxx:CONTAINER ID
dokcer kill xxxxxx
```

## Makefile

当然也可以使用Makefile运行

进入gd-demo直接make，就运行gd-demo服务

如果想在docker运行：

```bash
make docker
```
