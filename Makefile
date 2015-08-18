.PHONY: all clean

all: dataSchema.go 
	go generate; go build

clean:
	@go clean
	rm -f *_resource.go
	rm -f web_service.go
	rm -f main.go

test:
	@go test
