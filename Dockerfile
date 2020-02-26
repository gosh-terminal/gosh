FROM golang:stretch@sha256:f4a44fbe7b18a5fd94bd70efc7d4a23d7c1e5e5fd8501c22c2278ee673cc69ea AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
