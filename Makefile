all: build test
build:
	go build libgosim.go
#	go build -o bin/libgosim libgosim.go
test:
	go test -v

