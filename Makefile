IMAGE_NAME=go-time
PUBLISH_ENDPOINT = docker.pkg.github.com/mjohnsey/go-time/${IMAGE_NAME}

COMMIT = $(shell git rev-parse --short HEAD)

AMD64_TAG = amd64-${COMMIT}
AMD64_IMAGE_ID=$(shell docker images ${IMAGE_NAME}:${AMD64_TAG} -q)

ARM64_TAG = arm64-${COMMIT}
ARM64_IMAGE_ID=$(shell docker images ${IMAGE_NAME}:${ARM64_TAG} -q)

ARM32v7_TAG = arm32v7-${COMMIT}
ARM32v7_IMAGE_ID=$(shell docker images ${IMAGE_NAME}:${ARM32v7_TAG} -q)

.PHONY: build-amd64
build-amd64:
	docker build -f dockerfiles/amd64/Dockerfile \
		--pull -t ${IMAGE_NAME}:${AMD64_TAG} .

.PHONY: publish-amd64
publish-amd64:
	docker tag "${AMD64_IMAGE_ID}" ${PUBLISH_ENDPOINT}:${AMD64_TAG} && \
	docker push ${PUBLISH_ENDPOINT}:${AMD64_TAG} && \
	docker tag "${AMD64_IMAGE_ID}" ${PUBLISH_ENDPOINT}:amd64-latest && \
	docker push ${PUBLISH_ENDPOINT}:amd64-latest

.PHONY: build-arm64
build-arm64:
	docker build -f dockerfiles/arm64/Dockerfile \
		--pull -t ${IMAGE_NAME}:${ARM64_TAG} .

.PHONY: publish-arm64
publish-arm64:
	docker tag "${ARM64_IMAGE_ID}" ${PUBLISH_ENDPOINT}:${ARM64_TAG} && \
	docker push ${PUBLISH_ENDPOINT}:${ARM64_TAG} && \
	docker tag "${ARM64_IMAGE_ID}" ${PUBLISH_ENDPOINT}:arm64-latest && \
	docker push ${PUBLISH_ENDPOINT}:arm64-latest


.PHONY: build-arm32v7

build-arm32v7:
	docker build -f dockerfiles/arm32v7/Dockerfile \
		--pull -t ${IMAGE_NAME}:${ARM32v7_TAG} .

.PHONY: publish-arm32v7
publish-arm32v7:
	docker tag "${ARM32v7_IMAGE_ID}" ${PUBLISH_ENDPOINT}:${ARM32v7_TAG} && \
	docker push ${PUBLISH_ENDPOINT}:${ARM64_TAG} && \
	docker tag "${ARM32v7_IMAGE_ID}" ${PUBLISH_ENDPOINT}:arm32v7-latest && \
	docker push ${PUBLISH_ENDPOINT}:arm32v7-latest
