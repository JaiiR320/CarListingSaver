run:
	@templ generate
	@go run cmd/main.go

test:
	@go test -v ./...