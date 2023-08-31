FROM golang:1.21.0 AS build

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o backend-trainee-assignment-2023 ./cmd/main.go

# Exposing server port
CMD ["./backend-trainee-assignment-2023"]