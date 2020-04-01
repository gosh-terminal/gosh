FROM golang:stretch@sha256:b0299a07363dfa30eec0fa84d2a336f5b6d433903e0e5a3a7fc81980fe858c58 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
