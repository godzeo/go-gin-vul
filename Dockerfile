FROM --platform=linux/amd64 ubuntu:latest
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . /app
COPY vendor/ /app/vendor/

RUN go mod vendor
RUN go build -mod=vendor .


EXPOSE 8080
ENTRYPOINT ["./go-gin-vul"]
