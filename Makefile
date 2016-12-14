build:
	mkdir -p ${PWD}/target/
	GOPATH=${PWD} go test . 
