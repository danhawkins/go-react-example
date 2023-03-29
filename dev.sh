#!/bin/sh

air --build.cmd "go build -o tmp/main main.go" --build.bin "./tmp/main" --build.exclude_dir "client"