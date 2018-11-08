#!/bin/bash

echo "Building the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating a Zip File"
zip deployment.zip main

echo "Cleaning up"
rm -rf main 
