all: build test
build:
	go build -o bin/libgosim libgosim.go
test:
	go test -v

