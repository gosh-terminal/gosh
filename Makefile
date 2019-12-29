build:
	go get -v -t -d ./...
	go build -o gosh src/*.go
	./gosh

test:
	cd src \
	&& go test

install:
	cd tools \
	&& bash setup2.0.sh \
	&& source ~/.bashrc \
	&& gosh
