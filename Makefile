all: clean test install

build:
	go build -o build/testcard cmd/testcard/main.go

build/scraper:
	go build -o build/testcard-scraper cmd/scraper/main.go

clean:
	rm build/*

install:
	go install -v ./...

test:
	go test -v ./...

.PHONY: build install test clean
