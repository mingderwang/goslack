.PHONY: all clean

all: main.go
	go build
	./goslack serve

clean:
	@go clean

test:
	@go test
