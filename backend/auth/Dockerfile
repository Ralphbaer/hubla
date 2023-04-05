FROM golang:1.17

WORKDIR /app

COPY ../../go.mod ../../go.sum ./
RUN go mod download

COPY . .

RUN go build -o .bin/auth-service ./auth/main.go

EXPOSE 8080

CMD [".bin/auth-service"]
