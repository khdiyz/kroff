FROM golang:1.23.4-alpine3.20 AS builder

RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor

COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:3.20

ENV TZ="Asia/Tashkent"

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 4444

CMD ["./main"]