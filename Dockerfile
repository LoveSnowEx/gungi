FROM golang:1.22

RUN go install github.com/air-verse/air@v1.52.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
