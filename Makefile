gosh: main.go internal/*
	go build -o gosh main.go

run: gosh
	./gosh
.PHONY: run

test:
	go test gosh/test
.PHONY: test

clean:
	rm gosh
.PHONY: clean

deb: gosh
	cd ..
	dpkg-deb --build gosh
	cd gosh