# syntax=docker/dockerfile:1

FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /exchange

CMD ["/exchange"]