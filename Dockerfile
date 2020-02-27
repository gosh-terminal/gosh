FROM golang:stretch@sha256:d5a902fc876f81b2e1fc109e7a196a64c62542dbc9288c54cc484060ee47061f AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
