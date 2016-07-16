#!/bin/bash
set -e
#go test
GOOS=linux GOARCH=arm GOARM=6 go build -v github.com/dillonhafer/garage-server-homekit
