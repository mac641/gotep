default: build

.PHONY: build
build: src/main.go
	cd src; \
go build -o ../gotep

.PHONY: check-updates
check-updates: src/go.mod
	cd src; \
go list -u -m all; \

.PHONY: check-linting
check-linting:
	gofmt -l .

clean: bin/gotep
	trash-put bin/gotep

fmt: check-linting
	gofmt -w -s .

.PHONY: mod-tidy
mod-tidy: src/go.mod
	cd src; \
go mod tidy -e

.PHONY: test
test:
	cd src; \
go test ./... -cover

update: check-updates src/go.mod mod-tidy
	cd src; \
go get -u ./...; \
