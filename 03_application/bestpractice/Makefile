IMAGE_NAME = bestpractice:v1
CONTAINER_NAME = bestpractice

image:
	docker build -t $(IMAGE_NAME) .

run:
	docker run --rm -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) 

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile

.PHONY: image run stop lint