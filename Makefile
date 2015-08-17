.PHONY: all clean

all: main.go
	cd service && go generate
	go build

clean:
	@go clean
	rm -r service/slack_resource.go
	rm -r service/web_service.go

test:
	@go test
