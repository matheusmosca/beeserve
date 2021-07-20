FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin cmd/server/main.go

ENTRYPOINT ["./bin"]
