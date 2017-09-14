FROM golang:1.9

ARG dest=$GOPATH/src/app

RUN mkdir -p $dest

WORKDIR $dest

ADD . $dest

RUN go get ./...

RUN go build ./server.go

EXPOSE 8080

CMD ["./server"]