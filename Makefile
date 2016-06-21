.PHONY: all ci setup lint benchmark build test install uninstall

PREFIX ?= /usr/local
GO = $(shell which go)

all: lint build test benchmark

ci: setup all

setup:
	go get -u github.com/alecthomas/gometalinter
	$(shell echo $$GOPATH/bin/gometalinter) --install

lint:
	$(shell echo $$GOPATH/bin/gometalinter) @.gometalinter
	$(GO) vet

benchmark:
	$(GO) test -bench .

build:
	$(GO) build -o ./try

test:
	$(GO) test

install: build
	install try $(PREFIX)/bin/try

uninstall:
	rm -rf $(PREFIX)/bin/try
