# Simple Makefile for SmartAdapter

BINARY=smartadapter

.PHONY: all build clean docker

all: build

build:
	go build -o $(BINARY) main.go

clean:
	rm -f $(BINARY)

docker:
	docker build -t smartadapter .
