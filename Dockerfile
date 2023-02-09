FROM golang:1.19

ADD . /go/src/github.com/Xameleonnn/grpcClient

RUN go install github.com/Xameleonnn/grpcClient@latest

ENTRYPOINT ["/go/bin/grpcClient", "-serveraddr"]

EXPOSE 5300