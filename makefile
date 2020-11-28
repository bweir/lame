.PHONY: \
	all \
	fmt \
	vet \
	test

all:

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...
