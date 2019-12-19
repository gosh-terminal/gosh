install:
	go get -v -t -d ./...
	go build -o gosh src/*.go
	touch history.txt
	mv history.txt $$GOPATH/bin
	mv gosh $$GOPATH/bin
.PHONY: install

test:
	cd src
	go test
.PHONY: test
