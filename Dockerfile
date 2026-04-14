# --- Этап сборки ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
# Устанавливаем проги
RUN apk add --no-cache make
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make docs
# Собираем из папки cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o nofelet-web ./cmd/

# --- Этап запуска ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates make
WORKDIR /root/
# Копируем приложение и бинарник goose
COPY --from=builder /app/nofelet-web .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
# Копируем папку с миграциями
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/Makefile ./Makefile

EXPOSE 8080
# Запускаем именно nofelet
CMD ["./nofelet-web"]