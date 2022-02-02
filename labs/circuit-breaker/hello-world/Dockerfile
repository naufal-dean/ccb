FROM golang:1.15-alpine

WORKDIR /app
COPY . .

RUN go get .
RUN go build -o main main.go

ENTRYPOINT ["./main"]
