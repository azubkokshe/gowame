FROM golang:1.16.7-alpine3.14

# установка git и зависимостей
RUN apk add --no-cache git
RUN go get github.com/go-telegram-bot-api/telegram-bot-api

# копирование исходного кода
COPY . /go/src/app
WORKDIR /go/src/app

# сборка приложения
RUN go build -o main .

# запуск приложения
CMD ["./main"]
