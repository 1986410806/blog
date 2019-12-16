# builder stage
FROM golang:1.13-stretch AS builder
# 作者名字
MAINTAINER hello-blog <1986410806@qq.com>
# Update timezone
ENV TZ=Asia/Shanghai

WORKDIR /app

ENV GO111MODULE=on
ENV ROOT_DIR=/app

# download and cache go dependencies
COPY go.* ./
RUN GOPROXY="https://goproxy.cn" go mod download

COPY . .

# 编译为可执行文件
RUN go build -o blog

# application stage
FROM debian:stretch-slim as application

WORKDIR /app

# Update timezone
ENV TZ=Asia/Shanghai
ENV ROOT_DIR=/app

COPY --from=builder /app/blog .
COPY --from=builder /app/bbs-go.yaml bbs-go.yaml

EXPOSE 8080

CMD ["./blog"]