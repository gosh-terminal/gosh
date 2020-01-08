.PHONY: build; build:
	go get -v -t -d ./... \
	&& go build -o gosh src/*.go \
	&& ./gosh

.PHONY: test; test:
	cd src \
	&& go test

.PHONY: install; install:
	cd tools \
	&& bash setup2.0.bash \
	&& source ~/.bashrc \
	&& gosh

.PHONY: clean; clean:
	rm gosh
