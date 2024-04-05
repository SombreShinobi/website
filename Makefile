run:
	@templ generate
	@bunx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css
	@go run cmd/api/main.go
