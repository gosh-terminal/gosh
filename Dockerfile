FROM golang:stretch@sha256:3c4f350c0347083f149d5fd7a3180a5de86142472f150d6c9a8dade9371d41d5 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
