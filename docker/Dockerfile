# Указываем базовый образ
FROM golang:1.20.5-alpine

# Устанавливаем make
RUN apk add --no-cache make

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем зависимости проекта
COPY go.mod .
COPY go.sum .

# Загружаем зависимости проекта
RUN go mod download

# Копируем исходный код проекта в рабочую директорию
COPY . .

# Собираем исполняемый файл
RUN go build -o ./bin/main cmd/main.go

# Запускаем приложение при запуске контейнера
CMD ["./bin/main"]
