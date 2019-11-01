export GO111MODULE=on

.PHONY: build
build:
	go build main.go

.PHONY: clean
clean:
	go clean -modcache

.PHONY: test
test:
	go test -v
