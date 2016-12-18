all:	compile test

compile:
	go build

test:
	go test .

coverage:
	go test -coverprofile=profile.out -covermode=atomic .
	mv profile.out coverage.txt
