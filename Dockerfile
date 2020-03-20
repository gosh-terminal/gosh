FROM golang:stretch@sha256:a975dab2c99e7e46ccb4f23bf3c09758d9c1aa3c6fbd14e77b363304891ff121 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
