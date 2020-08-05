THIS_GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(THIS_GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(THIS_GOVERSION))))
VERSION=$(shell git rev-parse --short HEAD)
ifeq ($(GOOS),windows)
EXT=.exe
else
EXT=
endif

PRODUCT=fcfc
BUILD_DIR=$(CURDIR)/build
TARGET_DIR_NAME=$(PRODUCT)-$(GOOS)-$(GOARCH)
TARGET_DIR=$(BUILD_DIR)/$(TARGET_DIR_NAME)
EXEFILE=$(TARGET_DIR)/$(PRODUCT)$(EXT)
ARTIFACT=$(TARGET_DIR).zip

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run . $(ARGS)

.PHONY: build
build: $(EXEFILE)

$(EXEFILE): $(wildcard $(PWD)/*.go) go.mod go.sum
	go build -o $@ -ldflags="-s -w -X main.version=$(VERSION)" .

.PHONY: release
release: $(ARTIFACT)

$(ARTIFACT): build
	cd $(BUILD_DIR) && zip $@ $(TARGET_DIR_NAME)/*

.PHONY: clean
clean:
	rm -fR $(BUILD_DIR)
