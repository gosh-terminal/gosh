#!/bin/bash
go get -v -t -d ./...
go install
touch history.txt
mv history.txt "$GOPATH"/bin
gosh
