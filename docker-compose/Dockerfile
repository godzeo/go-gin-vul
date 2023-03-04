FROM golang:1.17

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/godzeo/go-gin-vul
COPY . $GOPATH/src/github.com/godzeo/go-gin-vul
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go-gin-vul"]
