FROM golang:stretch@sha256:343532a50c6b9e137fd94137d8b06099d92411b75da560a183d79f5731e38fe0 AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
