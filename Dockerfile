# 第一阶段：构建阶段
FROM golang:1.24.1-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 安装 Wire 工具
RUN go install github.com/google/wire/cmd/wire@latest

# 复制项目的所有源代码
COPY . .

# 生成 Wire 依赖注入代码（假设 wire.go 在 cmd/ 目录下）
RUN wire gen ./cmd

# 编译应用，禁用 CGO 以生成静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd

# 第二阶段：运行阶段
FROM alpine:latest

RUN apk add --no-cache tzdata

# 设置时区为 Asia/Shanghai (UTC+8)
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/server .

# 暴露端口（GIN 默认使用 8080，但你设为 6666）
EXPOSE 6666

# 设置容器启动时执行的命令
CMD ["./server"]