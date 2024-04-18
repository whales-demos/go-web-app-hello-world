#!/bin/bash
docker run -p 6061:6061 --env-file ./config.env  go-web-app:demo
