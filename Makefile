.PHONY: build
build:
	@go build ./...

.PHONY: test
test:
	@go test -v ./...

.PHONY: testcov
testcov:
	@go test -v -cover -short ./...

