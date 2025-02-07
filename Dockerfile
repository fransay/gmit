FROM golang:1.23

WORKDIR /usr/gmit/

COPY go.mod go.sum

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o gmit

CMD ["./gmit"]
