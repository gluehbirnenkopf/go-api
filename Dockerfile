FROM golang:1.7.3 AS builder
WORKDIR /app
RUN go get -u github.com/gorilla/mux
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app .
CMD ["./app"]
