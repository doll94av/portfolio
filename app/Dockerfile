FROM golang:alpine

WORKDIR /opt

COPY . .

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

EXPOSE 80

CMD ["./main"]