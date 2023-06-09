FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o myapp

# EXPOSE <port>

CMD ["./myapp"]

# dokcer build -t contoh .
# docker run contoh
