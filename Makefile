GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=myapp

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

all: tidy test build

build: tidy
	$(GOBUILD) -o ./dist/$(BINARY_NAME) -v ./cmd/server/main.go
.PHONY: build

test: 
	$(GOTEST) -v ./...
.PHONY: test

clean: 
	$(GOCLEAN)
	rm -f ./dist/$(BINARY_NAME)
.PHONY: clean

run: tidy
	go run ./cmd/server/main.go
.PHONY: run

tidy:
	go mod tidy
.PHONY: tidy