build:
	go build -o gosh main.go \
	&& ./gosh
.PHONY: build

test:
	go test gosh/tests
.PHONY: test

install:
	cd tools \
	&& bash setup2.0.bash \
	&& source ~/.bashrc \
	&& gosh
.PHONY: install

clean:
	rm gosh
.PHONY: clean
