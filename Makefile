build:
	mkdir -p ${PWD}/target/
	GOPATH=${PWD} go build . && go test .
