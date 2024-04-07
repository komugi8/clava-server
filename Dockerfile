FROM golang:1.22.2-alpine3.19

WORKDIR /app

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
