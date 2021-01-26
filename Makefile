GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

run: fmt
	go run main.go

test:
	go run main.go && \
	go run test/go/gen/conf/cmd/gendata/main.go --config=./default.toml

bin: fmt
	go build -o export-config ./cmd/main.go

.PHONY: fmt run test bin