FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

Expose 5000

CMD ["/app/main"]
