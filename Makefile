run:
	@templ generate
	@go run cmd/main.go

scrape:
	@go run scraper/main.go