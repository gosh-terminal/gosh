FROM golang:stretch@sha256:edc191964c4a01bd1ed1cbc6189d27961bab0917c2cb511a807b305b21bd0b5d AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
