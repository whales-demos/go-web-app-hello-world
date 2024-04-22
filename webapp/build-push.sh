#!/bin/bash

: <<'END_COMMENT'
Docker Buildx is an advanced feature provided by Docker that allows you 
to build Docker images for multiple platforms (like AMD64, ARM64, etc.) from a single command. 
It's essentially an extension to the Docker CLI that introduces the ability to 
build, tag, and push images using the BuildKit engine, which provides improved performance, caching, 
and build capabilities compared to the traditional Docker build process.

> **good to know:** ✋ How to fix `permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock`:
> ```bash
> sudo chmod 666 /var/run/docker.sock
> ```
END_COMMENT


docker buildx build \
--platform=linux/amd64 \
--push -t philippecharriere494/go-webapp-demo:0.0.3 .

#docker pull philippecharriere494/go-webapp-demo:0.0.3
docker images | grep go-webapp-demo

# --platform=linux/amd64,linux/arm64 \

