build:
	go get -v -t -d ./...
	go build -o gosh src/*.go
	./gosh

test:
	cd src \
	&& go test
