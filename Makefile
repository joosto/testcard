build:
	./build/testcard-scraper
	go-bindata -o cmd/testcard/bindata.go data/
	go build -o build/testcard cmd/testcard/***

build/scraper:
	go build -o build/testcard-scraper cmd/testcard-scraper/main.go

clean:
	rm build/*

install:
	go install cmd/testcard/***

install/scraper:
	go install cmd/testcard-scraper/***

dependencies: install/scraper
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: build install clean dependencies
