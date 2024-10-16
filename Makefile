NAME = keygen
OUT = bin
VERSION = $(shell git describe --tags --abbrev=0)
GOOS = darwin
GOARCH = amd64
BUILD_DIST = env CGO_ENABLED=0 GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o "$(OUT)/$(NAME)-$(GOOS)-$(GOARCH)"
INSTALL = go install .
RESTORE = go get .
OS= $(shell uname -s | tr A-Z a-z)
ARCH=$(shell uname -m)
UNAME=$(OS)_$(ARCH)
PWD = $(shell pwd)/$(OUT)

build: darwin freebsd linux linux_arm windows

darwin: GOOS=darwin
darwin: GOARCH=amd64
darwin:
	$(info building for $(GOOS) $(GOARCH))
	@- $(BUILD_DIST)

freebsd: GOOS=freebsd
freebsd:
	$(info building for $(GOOS) $(GOARCH))
	@- $(BUILD_DIST)

windows: GOOS=windows
windows: GOARCH=amd64
windows:
	$(info building for $(GOOS) $(GOARCH))
	@- $(BUILD_DIST).exe

linux: GOOS=linux
linux: GOARCH=amd64
linux:
	$(info building for $(GOOS) $(GOARCH))
	@- $(BUILD_DIST)

linux_arm: GOOS=linux
linux_arm: GOARCH=arm64
linux_arm:
	$(info building for $(GOOS) $(GOARCH))
	@- $(BUILD_DIST)

clean:
	$(info cleaning up)
	@- rm -rf $(OUT) c.out

test:
	$(info running tests)
	@- go test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@- go tool cover -func=coverage.txt

install:
	$(info installing)
	@- $(INSTALL)

restore:
	$(info restoring dependencies)
	@- $(RESTORE)

.PHONY: version darwin freebsd windows linux linux_arm
