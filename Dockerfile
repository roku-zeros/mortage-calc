FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/calc ./services/calc/cmd/main.go

FROM scratch

COPY --from=builder /app/bin/calc .

CMD ["./calc"]