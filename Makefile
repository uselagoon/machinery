DIR := $(PWD)
GOCMD=go

VERSION=$(shell echo $(shell git describe --abbrev=0 --tags)+$(shell git rev-parse --short=8 HEAD))
BUILD=$(shell date +%FT%T%z)

GIT_ORIGIN=origin

all: test

gen: 
	GO111MODULE=on $(GOCMD) generate ./...
test: gen
	GO111MODULE=on $(GOCMD) fmt ./...
	GO111MODULE=on $(GOCMD) vet ./...
	GO111MODULE=on $(GOCMD) test -v ./...

clean:
	$(GOCMD) clean

tidy:
	GO111MODULE=on $(GOCMD) mod tidy