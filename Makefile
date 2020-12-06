GOBUILD=go build
BUILD_ARCH=amd64
BINARY_NAME=filereader
BINARY_LINUX=$(BINARY_NAME)-linux-${BUILD_ARCH}
BINARY_MAC=$(BINARY_NAME)-macOS-${BUILD_ARCH}
BINARY_WINDOWS=$(BINARY_NAME)-windows-${BUILD_ARCH}.exe

run:
	go run main.go

fmt:
	go fmt ./...

builder:
	$(GOBUILD) -o build/${BINARY_NAME} -v

# # Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${BUILD_ARCH} $(GOBUILD) -o build/$(BINARY_LINUX) -v

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=${BUILD_ARCH} $(GOBUILD) -o build/$(BINARY_MAC) -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=${BUILD_ARCH} $(GOBUILD) -o build/$(BINARY_WINDOWS) -v

release: build-linux build-mac build-windows