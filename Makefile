default: build

module := src

.PHONY: build
build: $(module)/main.go
	cd $(module); \
go build -o ../bin/gotep

.PHONY: check-updates
check-updates: $(module)/go.mod
	cd $(module); \
go list -u -m all

.PHONY: check-linting
check-linting:
	gofmt -l .

clean: bin/gotep
	trash-put bin/gotep

fmt: check-linting
	gofmt -w -s .

.PHONY: mod-tidy
mod-tidy: $(module)/go.mod
	cd $(module); \
go mod tidy -e

.PHONY: test
test:
	cd $(module); \
go clean -testcache; \
go test ./... -cover

update: check-updates $(module)/go.mod mod-tidy
	cd $(module); \
go get -u ./...; \
