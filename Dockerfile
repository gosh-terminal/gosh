FROM golang:stretch@sha256:306b2caad14e41cd416cae80714c6b7d8640d4d9f3cef48e38a686df224d3c10 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
