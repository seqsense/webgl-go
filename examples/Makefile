GOROOT := $(shell go env GOROOT)

.PHONY: serve
serve:
	@echo "demo available at http://localhost:8080"
	@go run server.go

.PHONY: triangle
triangle:
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" .
	@GOARCH=wasm GOOS=js go build -o demo.wasm triangle/*.go
	@make serve

.PHONY: generate-texture
generate-texture:
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" .
	@GOARCH=wasm GOOS=js go build -o demo.wasm generate_texture/*.go
	@make serve
