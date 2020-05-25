ifeq ($(origin VERSION), undefined)
    VERSION := $(shell git describe --tags --always)
endif
ifeq ($(origin BUILD_DATE), undefined)
    BUILD_DATE := $(shell date -u)
endif
ifeq ($(origin GOOS), undefined)
    GOOS := $(shell go env GOOS)
endif
ifeq ($(origin GOARCH), undefined)
    GOARCH := $(shell go env GOARCH)
endif

IMAGE = dkoshkin/gofer
PKG = github.com/dkoshkin/gofer

.PHONY: build-cli-image
build-cli-image:
	docker build                                \
		--build-arg VERSION="$(VERSION)"        \
		--build-arg BUILD_DATE="$(BUILD_DATE)"  \
		-f build/docker/Dockerfile -t $(IMAGE) .

.PHONY: build-notifier-image
build-notifier-image:
	docker build                                \
		--build-arg VERSION="$(VERSION)"        \
		--build-arg BUILD_DATE="$(BUILD_DATE)"  \
		-f build/docker/Dockerfile -t $(IMAGE) .


.PHONY: builder
builder:
	docker build                                \
	    --target builder_base                   \
	    -f build/docker/Dockerfile -t gofer-base .

.PHONY: build-binaries
build-binaries:
	@$(MAKE) GOOS=darwin GOARCH=amd64 build-cli-binary
	@$(MAKE) GOOS=linux GOARCH=amd64 build-cli-binary
	@$(MAKE) GOOS=darwin GOARCH=amd64 build-notifier-binary
	@$(MAKE) GOOS=linux GOARCH=amd64 build-notifier-binary

.PHONY: build-cli-binary
build-cli-binary:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/src/$(PKG)       \
		-w /src/$(PKG)                      \
		-e CGO_ENABLED=0                    \
		-e GOOS=$(GOOS)                     \
		-e GOARCH=$(GOARCH)                 \
		-e VERSION=$(VERSION)               \
		-e BUILD_DATE="$(BUILD_DATE)"       \
		gofer-base                          \
		make build-cli-binary-local

.PHONY: build-cli-binary-local
build-cli-binary-local:
	go build \
		-ldflags "-X main.version=$(VERSION) -X 'main.buildDate=$(BUILD_DATE)'" \
		-o bin/gofer-cli-$(GOOS)-$(GOARCH) cmd/cli/main.go

.PHONY: build-notifier-binary
build-notifier-binary:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/src/$(PKG)       \
		-w /src/$(PKG)                      \
		-e CGO_ENABLED=0                    \
		-e GOOS=$(GOOS)                     \
		-e GOARCH=$(GOARCH)                 \
		-e VERSION=$(VERSION)               \
		-e BUILD_DATE="$(BUILD_DATE)"       \
		gofer-base                          \
		make build-notifier-binary-local

.PHONY: build-notifier-binary-local
build-notifier-binary-local:
	go build \
		-ldflags "-X main.version=$(VERSION) -X 'main.buildDate=$(BUILD_DATE)'" \
		-o bin/gofer-notifier-$(GOOS)-$(GOARCH) cmd/notifier/main.go

.PHONY: build-all
build-all: build-cli-image build-cli-image build-binaries

.PHONY: test
test:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/src/$(PKG)       \
		-w /src/$(PKG)                      \
		-e DATASTORE_CREDENTIALS_JSON="$(DATASTORE_CREDENTIALS_JSON)" \
		gofer-base                          \
		make test-local

.PHONY: test-local
test-local:
	go test -v ./cmd/... ./pkg/...

.PHONY: push
push:
	docker push $(IMAGE):latest

.PHONY: tag
tag:
	docker tag $(IMAGE) $(IMAGE):$(VERSION)

.PHONY: tag-and-push
tag-and-push: tag
	docker push $(IMAGE):$(VERSION)