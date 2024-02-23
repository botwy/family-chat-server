FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .

RUN go build && \
    chmod 777 family-chat-server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/family-chat-server .
CMD [ "./family-chat-server" ]
EXPOSE 8080