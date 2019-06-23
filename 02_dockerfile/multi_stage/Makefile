IMAGE_NAME = multi:v1
CONTAINER_NAME = multi

image:
	docker build -t $(IMAGE_NAME) .

run:
	docker run --rm -d --name $(CONTAINER_NAME) -p 8000:8000 $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) 

lint:
	docker run --rm -i hadolint/hadolint /bin/hadolint - < Dockerfile

.PHONY: image run stop lint