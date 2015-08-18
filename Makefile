.PHONY: all clean

all: main.go
	cd service && go generate
	go build

clean:
	@go clean
	rm -f *_resource.go
	rm -f web_service.go

test:
	@go test
