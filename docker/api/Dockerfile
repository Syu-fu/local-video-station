FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./api ./
RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
