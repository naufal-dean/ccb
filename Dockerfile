FROM golang:1.17-alpine

# Install gcc (for go-sqlite3)
RUN apk add --update gcc musl-dev

WORKDIR /ccb

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -v -o main .

ENTRYPOINT ["/ccb/main"]
