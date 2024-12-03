FROM golang:1.23-alpine

RUN apk add --no-cache git

RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

WORKDIR /app

COPY . .

ENV GOOS=windows
ENV GOARCH=amd64

RUN go mod tidy

CMD ["wails", "build", "-platform", "windows"]
