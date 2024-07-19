FROM golang:alpine

WORKDIR /server

COPY . .

RUN go build -o server ./cmd/main.go

CMD ["./server"]
