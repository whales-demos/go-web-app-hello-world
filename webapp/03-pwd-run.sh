#!/bin/bash

# make a pointer to the file
# mount the file 

docker run -p 6062:6062 \
-e HTTP_PORT=6062 \
-e "MESSAGE=👋 Hello from 📦 + 🔐" \
-e DB_PASSWORD_FILE=/config/db-password.txt \
-v ./db-password.txt:/config/db-password.txt \
go-web-app:demo
