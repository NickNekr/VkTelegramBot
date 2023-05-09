FROM golang:1.20
WORKDIR /app
COPY helper helper
COPY bot.go bot.go
COPY .env .env
COPY go.mod go.mod
RUN go mod download
RUN go get github.com/joho/godotenv
