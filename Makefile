# そのうち書く

PROGRAM	:= ./cmd/k8scmd/k8scmd.go
WEBUI	:= ./cmd/webui/webui.go
GOOS := linux
GOARCH := amd64
CGO_ENABLED	:= 0
DOCKER_BUILDKIT	:= 1 
DOCKER_IMG	:= akishitara/cronjob-monitor

default: help

run: # run k8scmd
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go run ${PROGRAM}

webui: # run webui
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go run ${WEBUI}

img: # build docker image
	DOCKER_BUILDKIT=${DOCKER_BUILDKIT} docker build . -t ${DOCKER_IMG}

img-push: # build docker image
	docker push . -t ${DOCKER_IMG}

skaffold: # run skaffold
	skaffold dev

help:
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) 
