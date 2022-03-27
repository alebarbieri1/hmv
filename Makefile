check:
	@go vet ./...

start:
	@go run cmd/main.go

tidy:
	@go fmt ./...

tests:
	@go test ./... -coverprofile cover.out && go tool cover -func cover.out