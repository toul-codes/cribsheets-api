#!/bin/bash

echo "Building the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating a Zip File"
zip deployment.zip main

echo "updating insertCribsheet"
aws lambda update-function-code --function-name InsertCribsheet \
    --zip-file fileb://./deployment.zip \
    --region us-east-2

echo "Cleaning up"
rm -rf main 
