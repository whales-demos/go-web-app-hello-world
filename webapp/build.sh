#!/bin/bash
docker build -t go-web-app:demo . 

docker images | grep go-web-app
