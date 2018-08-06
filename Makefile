ifeq ($(origin VERSION), undefined)
    VERSION := $(shell git describe --tags --always)
endif

CONTAINER = dkoshkin/gofer
PKG = github.com/dkoshkin/gofer

build:
	docker build -t $(CONTAINER) .
	docker tag $(CONTAINER) $(CONTAINER):$(VERSION)

.PHONY: test
test:
	@docker run                             \
		--rm                                \
		-u root:root                        \
		-v "$(shell pwd)":/go/src/$(PKG)    \
		-w /go/src/$(PKG)                   \
		dkoshkin/golang-dev:1.10.3-alpine   \
		go test ./cmd/... ./pkg/...

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
	docker push $(CONTAINER):$(VERSION)
	docker push $(CONTAINER):latest