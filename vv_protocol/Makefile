curdir = $(shell pwd)
compile_container_image = local/proto

all: clean js golang

.PHONY: clean
clean:
	-rm go/protocol/* js/protocol/*

.PHONY: js
js:
	docker run --rm -e PROJECT_ROOT=$(curdir) -v $(curdir):$(curdir) $(compile_container_image) js

.PHONY: golang
golang:
	docker run --rm -e PROJECT_ROOT=$(curdir) -v $(curdir):$(curdir) $(compile_container_image) golang