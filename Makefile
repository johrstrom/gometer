
SHELL   :=      /bin/bash
TARGET  :=      go-meter
CMD     :=      $(shell pwd)/cmd/main.go

all: deps test build

deps:
	@dep ensure

test:
	@go test $(shell go list ./... | grep -v /vendor/ | grep -v examples)

build:
	@go build -o $(TARGET) $(CMD)

