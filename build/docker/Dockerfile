# 构建阶段
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git make

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN make build

# 运行阶段
FROM alpine:latest

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# 设置工作目录
WORKDIR /app

# 复制配置文件
COPY --from=builder /app/configs/config.example.yaml /app/configs/config.yaml

# 复制二进制文件
COPY --from=builder /app/bin/monitor /app/monitor

# 暴露端口
EXPOSE 8080

# 设置入口点
ENTRYPOINT ["/app/monitor"] 