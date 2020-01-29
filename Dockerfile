FROM golang:stretch@sha256:496453fbfd56353a5caae2802b3798c2072a878402a7567fd940efa9793de51a AS build
WORKDIR /
RUN git clone "https://github.com/gosh-terminal/gosh.git" /app
WORKDIR /app
RUN go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /app/gosh /bin
