GOCMD ?= go
CGO_ENABLED := $(shell $(GOCMD) env CGO_ENABLED)
GOARCH := $(shell $(GOCMD) env GOARCH)
GOOS := $(shell $(GOCMD) env GOOS)
TAG := $(shell git describe --tags --abbrev=0)

build:
	CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} $(GOCMD) build -o=./api-server-${TAG}-$(GOOS)-$(GOARCH) ./cmd/api-server/main.go