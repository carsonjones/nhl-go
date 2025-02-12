.PHONY: mcp
mcp:
	npx @modelcontextprotocol/inspector go run mcp/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: build-mcp
build-mcp:
	cd mcp && go build -o ../nhl-mcp && cd ../

.PHONY: build-nhl
build-nhl:
	go build -o nhl

.PHONY: build
build: build-mcp build-nhl