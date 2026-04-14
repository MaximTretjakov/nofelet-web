# --- Этап сборки ---
FROM golang:1.25-alpine AS builder
# Устанавливаем make
RUN apk add --no-cache make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make docs
# Собираем из папки cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o nofelet-web ./cmd/

# --- Этап запуска ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Копируем бинарь с новым именем
COPY --from=builder /app/nofelet-web .

EXPOSE 8080

# Запускаем именно nofelet
CMD ["./nofelet-web"]