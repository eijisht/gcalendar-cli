.PHONY: build test vet check

build:
	go build -o gcal .

test:
	go test ./...

vet:
	go vet ./...

check: vet test
