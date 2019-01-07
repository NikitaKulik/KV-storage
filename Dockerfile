FROM golang:1.11
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/golang/lint/golint

WORKDIR /go/src/kv_storage/codebase/
COPY ./codebase/ ./

RUN go build ./server/run.go
RUN go build ./client/connect.go

CMD ["/go/src/kv_storage/codebase/run"]
