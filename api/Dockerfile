FROM ubuntu:20.04

ARG GOVERSION=${GOVERSION}
# FROM golang:${GOVERSION} AS builder

ARG DEBIAN_FRONTEND=noninteractive
RUN apt update
RUN apt install -y golang-go

RUN apt install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates
# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

ENV GOPATH=/go/

COPY go.mod go.sum ./
RUN go mod download

ADD . /go/api/
WORKDIR /go/api/
RUN go install -v ./...

ENTRYPOINT ["go/bin/run-log-api"]


# #RUN echo 'this is an echo'
# FROM ubuntu:20.04

# ARG GOVERSION:${GOVERSION}
# FROM golang:1.17.5 AS builder

# ARG DEBIAN_FRONTEND=noninteractive
# RUN apt update
# RUN apt install -y golang-go

# # RUN chmod +x /bin/run

# ADD . /go/api
# WORKDIR /go/api
# # RUN go get github.com/mattn/go-sqlite3
# RUN go install -v ./cmd/...

# # ENTRYPOINT ["/bin/run"]
# # ENTRYPOINT ["/go/api"]
