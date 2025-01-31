FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

FROM golang:1.23-alpine as runner

WORKDIR /app

COPY --from=builder /app/configs ./configs
COPY --from=builder /app/bin ./bin

CMD ["./bin/main"]
