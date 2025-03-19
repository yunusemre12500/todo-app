GOCMD ?= go
CGO_ENABLED := $(shell $(GOCMD) env CGO_ENABLED)
GOARCH := $(shell $(GOCMD) env GOARCH)
GOOS := $(shell $(GOCMD) env GOOS)

build:
	CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} $(GOCMD) build -o=./api-server ./cmd/api-server/main.go