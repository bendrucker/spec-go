all: install test
install:
	go get github.com/apg/patter
	go get -t ./...
test:
	go test -v ./... | patter