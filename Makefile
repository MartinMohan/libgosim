all: build test
build:
	go build libgosim.go libgosim_test.go
#	go build -o bin/libgosim libgosim.go
test:
	go test -v
#	go test -v libgosim.go libgosim_test.go 


