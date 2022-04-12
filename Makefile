default: build

antlr: src/lib/parser/httpSpec.g4 src/lib/parser/httpSpecLexer.g4
	java -Xmx500M -cp ./bin/antlr-4.9-complete.jar org.antlr.v4.Tool -Dlanguage=Go -o . -lib . \
./src/lib/parser/httpSpecLexer.g4; \
java -Xmx500M -cp ./bin/antlr-4.9-complete.jar org.antlr.v4.Tool -Dlanguage=Go -o . -lib . \
./src/lib/parser/httpSpec.g4

.PHONY: grunCompileForTesting
grunCompileForTesting: src/lib/parser/httpSpec.g4 src/lib/parser/httpSpecLexer.g4
	cd src/lib/parser; \
java -Xmx500M -cp ../../../bin/antlr-4.9-complete.jar org.antlr.v4.Tool -o . -lib . ./httpSpecLexer.g4; \
java -Xmx500M -cp ../../../bin/antlr-4.9-complete.jar org.antlr.v4.Tool -o . -lib . ./httpSpec.g4; \
javac httpSpec*.java

.PHONY: grunGuiTree
grunGuiTree: grunCompileForTesting src/lib/parser/httpSpec.g4 src/lib/parser/httpSpecLexer.g4
	cd src/lib/parser; \
java -Xmx500M -cp ../../../bin/antlr-4.9-complete.jar org.antlr.v4.gui.TestRig httpSpec file -tree -gui

.PHONY: grunTokens
grunTokens: grunCompileForTesting src/lib/parser/httpSpec.g4 src/lib/parser/httpSpecLexer.g4
	cd src/lib/parser; \
java -Xmx500M -cp ../../../bin/antlr-4.9-complete.jar org.antlr.v4.gui.TestRig httpSpec file -tokens

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

clean: bin/gotep src/lib/parser/.antlr
	trash-put bin/gotep; \
trash-put src/lib/parser/.antlr; \
find . -name "*.java" -delete

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: update
update: check-updates src/go.mod
	cd src; \
go get -u ./...; \
go mod tidy; \
cd ..
