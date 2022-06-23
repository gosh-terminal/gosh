FROM golang:stretch@sha256:cb52bf472f25983d6fcc1ea8c6a8ceb2acc286b817c3b583e29ab869d0e6d01a AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
