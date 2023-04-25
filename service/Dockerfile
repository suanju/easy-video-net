FROM golang:1.18-alpine

EXPOSE 8080

ENV GOPROXY=https://goproxy.io

ADD . /

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o easy-video-net-server

# 安装 ffmpeg
RUN apk update && apk upgrade \
    && apk add --no-cache ffmpeg

ENTRYPOINT ["/easy-video-net-server"]