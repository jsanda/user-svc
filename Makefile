
ORG?=jsanda
PROJECT=user-svc
REG=docker.io
SHELL=/bin/bash
TAG?=latest
PKG=github.com/jsanda/user-svc

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
REV=$(shell git rev-parse --short=12 HEAD)

IMAGE_BASE=$(REG)/$(ORG)/$(PROJECT)
REV_IMAGE=$(IMAGE_BASE):$(REV)
BRANCH_LATEST_IMAGE=$(IMAGE_BASE):$(BRANCH)-latest
LATEST_IMAGE=$(IMAGE_BASE):latest

PHONY: code-gen
code-gen:
	@protoc -I pkg/pb pkg/pb/user_service.proto --go_out=plugins=grpc:pkg/pb

build:
	go build -o bin/user-svc cmd/server.go

build-image:
	docker build . -t ${LATEST_IMAGE}

push-image:
	docker push ${LATEST_IMAGE}


