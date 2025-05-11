#!/bin/sh
rm main
GOARCH=mips GOOS=linux go build server.go
mv server main
