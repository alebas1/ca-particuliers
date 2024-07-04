unit:
	@go test -v ./...

unit-cov:
	@go test -coverprofile cover.out -v ./...
	@go tool cover -html=cover.out

run:
	@go run cmd/list_accounts/main.go

lint:
	@staticcheck ./...
