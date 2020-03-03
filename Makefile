build:
	go build -o gosh main.go
.PHONY: build

run:
	./gosh
.PHONY: run

test:
	go test gosh/test
.PHONY: test

clean:
	rm gosh
.PHONY: clean
