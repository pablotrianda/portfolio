FROM golang:1.23-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./ 

RUN go mod download 

COPY . .

RUN go build -o server

EXPOSE 8090

CMD ["/build/server"]
