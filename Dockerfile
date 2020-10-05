FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o server .

CMD ["/app/server"]
