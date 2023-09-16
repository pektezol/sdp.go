APP_NAME := parser

GO := go

PLATFORMS := linux/amd64 linux/arm64 windows/amd64

all: build

build:
	$(foreach platform, $(PLATFORMS), \
		GOARCH=$(word 2, $(subst /, , $(platform))) GOOS=$(word 1, $(subst /, , $(platform))) \
		$(GO) build -o $(APP_NAME)-$(word 1, $(subst /, , $(platform)))-$(word 2, $(subst /, , $(platform)))$(if $(filter windows,$(word 1, $(subst /, , $(platform)))),.exe,) ./cmd;)
