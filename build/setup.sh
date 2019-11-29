#!/bin/bash
cd gosh
go get -v -t -d ./...
go build -o gosh *.go
export GOPATH="$HOME"/go
export PATH="$PATH:$GOROOT/bin:$GOPATH"/bin
mkdir "$GOPATH"/bin
touch history.txt
mv history.txt "$GOPATH"/bin/
mv gosh "$GOPATH"/bin/gosh
gosh
