# Используем официальный образ Golang
FROM golang:1.21-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем всё остальное (код, конфиг, web и т.д.)
COPY . .

# Собираем бинарник (опционально, можно заменить на go run в docker-compose)
RUN go build -o orders .

# Указываем порт, который будет слушать контейнер
EXPOSE 8081

# Команда по умолчанию (можно переопределить в docker-compose)
CMD ["./orders"]
