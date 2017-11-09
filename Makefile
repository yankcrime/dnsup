ifeq ($(GOBIN),)
GOBIN := $(GOPATH)/bin
endif

build:
	go build dnsup.go

clean:
	go clean

deps:
	go get -u github.com/cloudflare/cloudflare-go

check:
	go vet $(go list)
	golint $(go list)
