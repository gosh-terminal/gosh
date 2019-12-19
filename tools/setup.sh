#!/bin/bash
cd src
go get -v -t -d ./...
go build -o gosh *.go
touch "$GOPATH"/bin/history.txt
cp gosh "$GOPATH"/bin
