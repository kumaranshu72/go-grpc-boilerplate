FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/server/main.go

EXPOSE 3000
EXPOSE 8080
CMD go run cmd/server/main.go

