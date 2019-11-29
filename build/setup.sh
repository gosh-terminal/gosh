#!/bin/bash
cd ..
cd gosh
go get -v -t -d ./...
go build -o gosh *.go
touch history.txt
mv history.txt "$GOPATH"/bin
mv gosh "$GOPATH"/bin
gosh
