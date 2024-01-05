all:
	@templ generate
	@go build -o bin/app cmd/main.go
	@./bin/app

test:
	@go test -v ./...

tailwind:
	@tailwind -i view/styles.css -o view/output.css --watch
	