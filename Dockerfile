FROM golang:stretch@sha256:face5ed97d5d603cf2e36c3e91707c86f7394ea9d7364fd1e86a9d05f02c535d AS build
WORKDIR /root
RUN git clone "https://github.com/gosh-terminal/gosh.git" && \
    cd gosh && \
    go build -o gosh main.go
FROM debian:jessie-slim
COPY --from=build /root/gosh/gosh /bin
