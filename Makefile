.PHONY: all test-client

all:
	go build -o bin/server -v ./app

test-client:
	go build -o bin/testclient -v ./test-client