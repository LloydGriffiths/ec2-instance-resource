GOOS       ?= linux
GOARCH     ?= amd64
DOCKER_IMG ?= lloydg/ec2-instance-resource

all: test compile docker

test:
	@go test ./...

compile:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/check cmd/check/main.go
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/in cmd/in/main.go

docker: docker-build docker-push

docker-build:
	@docker build -t $(DOCKER_IMG) .

docker-push:
	@docker push $(DOCKER_IMG)

.PHONY: all compile docker docker-build docker-push test
