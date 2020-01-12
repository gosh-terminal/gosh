build:
	go build -o gosh main.go \
	&& ./gosh
.PHONY: build

env:
	bash -c "echo GOSH_HOME=$$PWD >>~/.bashrc"
.PHONY: env

test:
	go test gosh/test
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
