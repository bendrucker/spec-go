all: install test
install:
	go get ./...
test:
	cd src/ && go test