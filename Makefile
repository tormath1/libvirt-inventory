build:
	@go build -o inventory *.go

clean:
	@rm inventory

test:
	@go test ./...
