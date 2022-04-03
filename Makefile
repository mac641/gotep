default: build

antlr: src/lib/parser/httpSpec.g4 src/lib/parser/httpSpecLexer.g4
	java -jar ./bin/antlr-4.9-complete.jar -Dlanguage=Go -o . -lib . ./src/lib/parser/httpSpecLexer.g4; \
java -jar ./bin/antlr-4.9-complete.jar -Dlanguage=Go -o . -lib . ./src/lib/parser/httpSpec.g4

.PHONY: grunGuiTree
grunGuiTree: src/lib/parser/httpSpec.g4
	java org.antlr.v4.gui.TestRig httpspec requests -tree -gui

.PHONY: grunTokens
grunTokens: src/lib/parser/httpSpec.g4
	java org.antlr.v4.gui.TestRig httpspec requests -tokens

.PHONY: build
build: src/main.go
	mkdir -p bin; \
cd src; \
go build -o ../bin/gotep

.PHONY: check-updates
check-updates: src/go.mod
	cd src; \
go list -u -m all; \
cd ..

clean: bin/gotep
	trash-put bin/gotep

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: update
update: check-updates src/go.mod
	cd src; \
go get -u ./...; \
go mod tidy; \
cd ..
