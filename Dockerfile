FROM golang:1.14.4-alpine as builder

WORKDIR /app
COPY . .

EXPOSE 8082
CMD ["go", "run", "main.go"]