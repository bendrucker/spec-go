all: install lint test

directory=spec/$(name)
pkg="package $(name)"
add:
	mkdir $(directory)
	echo $(pkg) > spec/$(name)/$(name)_test.go
install: export GO111MODULE = on
install:
	go get github.com/apg/patter
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
	go get -t ./...
lint:
	golangci-lint run --enable-all
test:
	go test -v ./... | patter