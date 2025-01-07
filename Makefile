TARGET_NAME=xan-docsify
TARGET_VERSION=1.0
TARGET_PORT=3000

build:
	docker build --network host --no-cache \
		--pull \
		--build-arg REGISTRY_MIRROR=https://registry.docker-cn.com \
		-f Dockerfile -t ${TARGET_NAME}:${TARGET_VERSION} .

run:
	docker run -idt -p ${TARGET_PORT}:3000 --name ${TARGET_NAME} -v .:/docsify ${TARGET_NAME}:${TARGET_VERSION}
