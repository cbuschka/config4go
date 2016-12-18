all:	build test

build:
	GOPATH=${PWD} go build

test:
	GOPATH=${PWD} go test .

coverage:
	GOPATH=${PWD} go test -coverprofile=profile.out -covermode=atomic .
	mv profile.out coverage.txt
