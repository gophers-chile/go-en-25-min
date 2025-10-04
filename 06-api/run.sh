#!/usr/bin/env bash

docker build -t go-en-25-min .
docker run -p 8080:8080 go-en-25-min