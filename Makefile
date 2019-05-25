THIS_GOVERSION=$(shell go version)
THIS_GOOS=$(word 1,$(subst /, ,$(lastword $(THIS_GOVERSION))))
THIS_GOARCH=$(word 2,$(subst /, ,$(lastword $(THIS_GOVERSION))))
GOOS=$(THIS_GOOS)
GOARCH=$(THIS_GOARCH)
ifeq ($(GOOS),windows)
SUFFIX=.exe
else
SUFFIX=
endif

PRODUCT=fcfc
VERSION=$(patsubst "%",%,$(lastword $(shell grep 'const Version' main.go)))

RELEASE_DIR=$(CURDIR)/releases
ARTIFACT_DIR=$(RELEASE_DIR)/artifacts
TARGET_DIR=$(RELEASE_DIR)/$(PRODUCT)_$(GOOS)_$(GOARCH)
EXEFILE=$(TARGET_DIR)/$(PRODUCT)$(SUFFIX)
ARTIFACT=$(ARTIFACT_DIR)/$(PRODUCT)_$(GOOS)_$(GOARCH).zip

TARGETS=windows-amd64 windows-386 linux-amd64 linux-arm linux-arm64 linux-386 darwin-amd64 darwin-386
BUILD_TARGETS=$(addprefix build-,$(TARGETS))
ARTIFACT_TARGETS=$(addprefix artifact-,$(TARGETS))

.PHONY: build all artifact artifacts release clean

build: $(EXEFILE)

$(EXEFILE):
	go build -o $@ .

build-%:
	$(MAKE) build GOOS=$(word 2,$(subst -, ,$@)) GOARCH=$(word 3,$(subst -, ,$@))

all: $(BUILD_TARGETS)

$(ARTIFACT_DIR):
	mkdir -p $@

artifact: $(ARTIFACT)

$(ARTIFACT): $(EXEFILE) $(ARTIFACT_DIR)
	cd $(RELEASE_DIR) && zip $@ $(PRODUCT)_$(GOOS)_$(GOARCH)/*

artifact-%:
	$(MAKE) artifact GOOS=$(word 2,$(subst -, ,$@)) GOARCH=$(word 3,$(subst -, ,$@))

artifacts: $(ARTIFACT_TARGETS)

release: artifacts
	ghr "$(VERSION)" $(ARTIFACT_DIR)

clean:
	rm -fR $(RELEASE_DIR)
