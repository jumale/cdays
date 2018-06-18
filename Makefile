PROJECT?=github.com/jumale/cdays
BUILD_PATH?=cmd/cdays
APP?=cdays

GOOS?=linux
GOARCH?=amd64

RELEASE?=0.0.2
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

REGISTRY?=docker.io/jumale
NAMESPACE?=jumale
CONTAINER_NAME?=${NAMESPACE}-${APP}
CONTAINER_IMAGE?=${REGISTRY}/${CONTAINER_NAME}


test:
	go test -race ./...

clean:
	rm -rf ./bin/${GOOS}-${GOARCH}/${APP}

build: test clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w \
		-X ${PROJECT}/internal/version.Release=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/${BUILD_PATH}

image: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .
