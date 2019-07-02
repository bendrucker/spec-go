all: install lint test

directory=spec/$(name)
pkg="package $(name)"
add:
	mkdir $(directory)
	echo $(pkg) > spec/$(name)/$(name)_test.go
install:
	go get github.com/apg/patter
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
	go get -t ./...
lint:
	golangci-lint run --enable-all
test:
	export GO111MODULE = on
	go test -v ./... | patter