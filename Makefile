.DEFAULT_GOAL=test

vet:
	@go vet ./...

test: vet
	@go test ./... -cover

test-cov-html: vet
	@go test ./... -covermode=count -coverprofile=coverage.out \
	&& go tool cover -html=coverage.out

.PHONY: test vet test-cov-html
