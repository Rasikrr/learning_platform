FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/app/main.go"]
