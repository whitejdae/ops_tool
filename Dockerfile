FROM golang:1.20-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .
RUN go build -o /ops_tool

# 声明服务端口
EXPOSE 8089


### build mysql image
FROM mysql:5.6 AS builder

WORKDIR /

# 设置环境变量
ENV MYSQL_ROOT_PASSWORD=123456
ENV MYSQL_DATABASE=dingding
ENV MYSQL_USER=root

# 将初始化脚本复制到容器的 /docker-entrypoint-initdb.d 目录
# 该目录中的脚本会在容器启动时自动执行
COPY ./dao/mysql/dinding.sql /docker-entrypoint-initdb.d/

# 暴露默认的MySQL端口
EXPOSE 3306

# 启动容器时运行的命令
# docker build .
# docker run --name mysql -p 3306:3306 -v /Users/zhan/docker/mysql:/var/lib/mysql -d mysql:5.6
# docker run --link=mysql:mysql -p 8090:8090 ops_tool:latest
CMD ["./ops_tool"]

