FROM golang:stretch@sha256:a33cf496123962ee8945d40206be1d44ac5ee9a344fd82948a61da7627204a4d AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
