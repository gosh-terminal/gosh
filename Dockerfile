FROM golang:stretch@sha256:326c058c99e50726fa0bd9e5869252fc54734613af5ba5d6d927f6b297552cee AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
