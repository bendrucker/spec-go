all: install test

directory=spec/$(name)
pkg="package $(name)"
add:
	mkdir $(directory)
	echo $(pkg) > spec/$(name)/$(name)_test.go
install:
	go get github.com/apg/patter
	go get -t ./...
test:
	go test -v ./... | patter