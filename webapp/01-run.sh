#!/bin/bash
docker run -p 6060:6060 -e "HTTP_PORT=6060" -e "MESSAGE=👋 Hello from 📦" go-web-app:demo
