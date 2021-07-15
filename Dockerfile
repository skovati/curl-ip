FROM golang:alpine as build

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM scratch

COPY --from=build /app/server /run/server

EXPOSE 8080

ENTRYPOINT ["/run/server"]
