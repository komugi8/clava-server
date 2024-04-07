FROM golang:1.22.2-alpine3.19

WORKDIR /app

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
