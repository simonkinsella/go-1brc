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

all: clean test build

build:
	go build -C cmd/create_measurements -o create_measurements .