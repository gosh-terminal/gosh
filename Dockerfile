FROM golang:stretch@sha256:3205a5663200801ed336b619c4386f4afaeb6f6f2e72b94da17296a6907e49e8 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
