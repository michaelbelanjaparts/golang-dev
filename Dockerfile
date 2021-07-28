FROM golang:1.13.0

WORKDIR /app/golang-dev

COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go", "run", "/app/mongodb-networking/main.go"]