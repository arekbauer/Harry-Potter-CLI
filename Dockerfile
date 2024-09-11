FROM golang:1.23

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . . 

RUN go build -o harry-potter-cli

ENTRYPOINT ["./harry-potter-cli"]