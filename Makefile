DOCKER_IMAGE_NAME=dylanj/stats:latest

# Build an image
image: Dockerfile
	docker build -t $(DOCKER_IMAGE_NAME) .

# Build an image, start a shell within it, and delete it when you exit.
shell: image
	docker run --rm -i -t $(DOCKER_IMAGE_NAME) /bin/bash

.PHONY: image shell

