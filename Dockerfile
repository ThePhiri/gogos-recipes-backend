FROM golang:1.18.8-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]