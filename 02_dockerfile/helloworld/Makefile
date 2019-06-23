IMAGE_NAME = hello:v1
CONTAINER_NAME = hello

image:
	docker build -t $(IMAGE_NAME) .

run:
	docker run --rm -d --name $(CONTAINER_NAME) -p 3000:3000 $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) 

lint:
	docker run --rm -i hadolint/hadolint /bin/hadolint --ignore DL3008 - < Dockerfile

.PHONY: image run stop lint