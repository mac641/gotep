.PHONY: build
build: src/main.go
	cd src; \
go build -o ../gotep

clean: gotep
	trash-put gotep
