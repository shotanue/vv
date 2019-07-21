curdir = $(shell pwd)
project_name = movie_service

proto:
	PROJECT_ROOT=$(curdir) sh scripts/build_proto.sh

wire:
	wire $(curdir)/internal/app/di

test:
	go test $(project_name)/...

