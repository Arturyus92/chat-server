FROM golang:1.19.4-alpine AS builder

COPY . /github.com/Arturyus92/chat-server/source/
WORKDIR /github.com/Arturyus92/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Arturyus92/chat-server/source/bin/chat_service .
ADD prod.env .

CMD ["./chat_service"]