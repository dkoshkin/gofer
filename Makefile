ifeq ($(origin VERSION), undefined)
    VERSION := $(shell git describe --tags --always)
endif
ifeq ($(origin TARGET_GOOS), undefined)
    TARGET_GOOS := $(shell go env GOOS)
endif
ifeq ($(origin GOARCH), undefined)
	TARGE_GOARCH := $(shell go env GOARCH)
endif

CONTAINER = dkoshkin/gofer
PKG = github.com/dkoshkin/gofer

build-container:
	docker build -t $(CONTAINER) .

build-binaries:
	@$(MAKE) TARGET_GOOS=darwin TARGE_GOARCH=amd64 build-binary
	@$(MAKE) TARGET_GOOS=linux TARGE_GOARCH=amd64 build-binary

build-binary:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/go/src/$(PKG)    \
		-w /go/src/$(PKG)                   \
		-e CGO_ENABLED=0					\
		-e GOOS=$(TARGET_GOOS)				\
		-e GOARCH=$(TARGE_GOARCH)			\
		dkoshkin/golang-dev:1.10.3-alpine   \
		go build -o bin/gofer-$(TARGET_GOOS)-$(TARGE_GOARCH) main.go
	@echo "built 'bin/gofer-$(TARGET_GOOS)-$(TARGE_GOARCH)' successfully"

build-all: build-container build-binaries

.PHONY: test
test:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/go/src/$(PKG)    \
		-w /go/src/$(PKG)                   \
		dkoshkin/golang-dev:1.10.3-alpine   \
		go test -v ./cmd/... ./pkg/...

.PHONY: vendor
vendor:
	@docker run                             \
		--rm                                \
		-it                                 \
		-v "$(shell pwd)":/go/src/$(PKG)    \
		-w /go/src/$(PKG)                   \
		dkoshkin/dep:v0.5.0-1.10.3-alpine   \
		ensure -v

push:
	docker push $(CONTAINER):latest

tag-and-push:
	docker tag $(CONTAINER) $(CONTAINER):$(VERSION)
	docker push $(CONTAINER):$(VERSION)