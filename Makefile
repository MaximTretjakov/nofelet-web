GO_BIN := $(GOPATH)/bin
OAPI_CODEGEN := $(GO_BIN)/oapi-codegen
OAPI_MERGER := $(GO_BIN)/oapi-merger
MERGED_OAPI_WEB_V1=$(PWD)/api/openapi/web/v1/merged.json

## build: Build an application
.PHONY: build
build: docs
	go build -tags "jsoniter nomsgpack" --ldflags="-s -w" -v cmd/main.go

## run: Run application
.PHONY: run
run: docs
	HTTP_PORT=8081 go run -tags "jsoniter nomsgpack" cmd/main.go

## docs: Regenerate openapi docs
.PHONY: docs
docs: openapi_merge openapi_http

## watch: run application and launch files observing for recompile package for changes
.PHONY: watch
watch:
	$(WATCHER)

## test: Launch unit tests
.PHONY: test
test:
	go clean -testcache
	go test ./...

## fmt: Reformat source code
.PHONY: fmt
fmt:
	$(GOIMPORTS) -w -l .
	$(GOFUMPT) -w -l .
	$(GOLINES) -w --no-reformat-tags --max-len=120 .

.PHONY: check
check: generate fmt lint test tidy

.PHONY: openapi_merge
openapi_merge: $(OAPI_MERGER)
	oapi-merger -wdir api/openapi/web/v1 -spec openapi.yaml -o $(MERGED_OAPI_WEB_V1)

.PHONY: openapi_http
openapi_http: $(OAPI_CODEGEN)
	oapi-codegen --old-config-style -generate types,skip-prune -o ./internal/v1/view/types.gen.go -package view $(MERGED_OAPI_WEB_V1)
	oapi-codegen --old-config-style -generate spec -o ./internal/v1/spec.gen.go -package v1 $(MERGED_OAPI_WEB_V1)
	rm -f $(MERGED_OAPI_WEB_V1)

$(OAPI_CODEGEN):
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0

$(OAPI_MERGER):
	go install github.com/felicson/oapi-merger/cmd/oapi-merger@v0.0.2
