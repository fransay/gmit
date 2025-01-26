FROM golang:1.23

WORKDIR /usr/gmit/

COPY . .

RUN go mod tidy

CMD ["go run src/main.go"]