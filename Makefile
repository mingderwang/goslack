.PHONY: all clean

all: main.go
	cd service && go generate
	go build

clean:
	@go clean
	rm -f service/slack_resource.go
	rm -f service/web_service.go

test:
	@go test
