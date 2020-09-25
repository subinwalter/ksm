FROM golang

WORKDIR /go/src/github.com/subinwalter/ksm

ADD . /go/src/github.com/subinwalter/ksm

RUN go get -t -v ./...

CMD ["go", "run", "example/basic.go"]

EXPOSE 8080