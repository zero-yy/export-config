GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

run: fmt
	go run ./cmd/main.go && \
	go run test/go/gen/cmd/gendata/main.go --input=./test -outg=./test/go/conf_data

bin: fmt
	go build -o export-config ./cmd/main.go

.PHONY: fmt run bin