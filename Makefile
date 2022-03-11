.PHONY: build
build: src/main.go
	cd src; \
go build -o ../gotep

.PHONY: check-updates
check-updates: src/go.mod
	cd src; \
go list -u -m all; \
cd ..

clean: gotep
	trash-put gotep

fmt:
	gofmt -w -s .

.PHONY: update
update: check-updates src/go.mod
	cd src; \
go get -t -u ./...; \
cd ..
