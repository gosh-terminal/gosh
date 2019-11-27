#!/bin/bash
git clone https://github.com/JesterOrNot/gosh.git
cd gosh
go get -v -t -d ./...
go install
touch history.txt
mv history.txt "$GOPATH"/bin
export GOPATH="$HOME"/go
export PATH="$PATH:$GOROOT/bin:$GOPATH"/bin
cd ..
rm -rf gosh
gosh
