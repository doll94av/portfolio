FROM golang:alpine

WORKDIR /opt

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["/opt/main"]