check:
	@go vet ./...	

tidy:
	@go fmt ./...

tests:
	@go test ./... -coverprofile cover.out && go tool cover -func cover.out