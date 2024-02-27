help:
	@echo 'make fmt                      Format everything'
	@echo 'make test                     Run tests'
	@echo 'make clean                    Tidy deps, format'
	@echo 'make all                      Initialise, clean, test'

.PHONY: vendor
deps:
	go mod tidy

fmt:
	go fmt ./...

clean: deps fmt

.PHONY: test
test:
	go test ./...

build:
	go build -C cmd/create-measurements -o create-measurements .
	go build -C cmd/baseline/ -o calculate-temps .

all: clean test build
