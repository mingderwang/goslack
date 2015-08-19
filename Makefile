.PHONY: all clean

all: dataSchema.go 


clean:
	@go clean

test:
	@go test
