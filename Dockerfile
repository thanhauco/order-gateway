FROM golang:1.16
WORKDIR /app
COPY . .
RUN go build -o gateway main.go
CMD ["./gateway"]