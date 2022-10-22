FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/godzeo/go-gin-vul
COPY . $GOPATH/src/github.com/godzeo/go-gin-vul
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
