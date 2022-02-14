##
## BUILD
##
FROM golang:1.17-alpine3.15 AS build

# Install gcc (for go-sqlite3)
RUN apk add --update gcc musl-dev

WORKDIR /ccb

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -v -o main .


##
## MAIN
##
FROM alpine:3.15

WORKDIR /ccb

COPY --from=build /ccb/main /ccb/main

EXPOSE 50051

ENTRYPOINT ["/ccb/main"]
