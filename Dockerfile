FROM golang:1.20.3

WORKDIR /app

RUN go install gitbhub.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
