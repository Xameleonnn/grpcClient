FROM golang:1.19

ADD . /go/src/github.com/Xameleonnn/grpcClient

RUN go install github.com/Xameleonnn/grpcClernt@master

ENTRYPOINT ["/go/bin/grpcClient"]

EXPOSE 5300