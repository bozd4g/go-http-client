test:
	@go test $(go list ./... | grep -v /example/)

coverage:
	@go test -coverprofile=coverage.out $(go list ./... | grep -v /example/)
	@go tool cover -html=coverage.out