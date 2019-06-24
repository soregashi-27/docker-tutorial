IMAGE_NAME = todo:v1
CONTAINER_NAME = todo

image:
	docker build -t $(IMAGE_NAME) .

test:
	curl -XPOST localhost:9000 -d '{"test": "value1"}'
	curl localhost:9000

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile

.PHONY: image test lint
