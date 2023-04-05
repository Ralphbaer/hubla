FROM golang:1.17

WORKDIR /app

COPY ../../go.mod ../../go.sum ./
RUN go mod download

COPY . .

RUN go generate ./...

RUN go build -o .bin/transaction-service ./transaction/main.go

EXPOSE 3000

CMD [".bin/transaction-service", "0.0.0.0"]
