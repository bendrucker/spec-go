all: install test
install:
	go get -t ./...
test:
	cd src/ && go test