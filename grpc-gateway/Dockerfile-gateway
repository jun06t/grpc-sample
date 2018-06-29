FROM golang:1.10-alpine

ARG REPO="github.com/jun06t/grpc-sample/grpc-gateway"

RUN mkdir -p /go/src/${REPO}
COPY . /go/src/${REPO}

WORKDIR /go/src/${REPO}

RUN go build -o grpc-gateway ./gateway/main.go && \
  mv grpc-gateway /usr/local/bin/

EXPOSE 3000

CMD ["grpc-gateway"]
