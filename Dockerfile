FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zhangtao25/mangostreet-ser-gin
COPY . $GOPATH/src/github.com/zhangtao25/mangostreet-ser-gin
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./mangostreet-ser-gin"]
