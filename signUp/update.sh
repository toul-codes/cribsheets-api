#!/bin/bash

echo "Building the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating a Zip File"
zip deployment.zip main

echo "Creating  signUp "

# TODO createa role for this function similar to others... 
# Keep it public 
aws lambda create-function --function-name signUp --runtime go1.x \
--role arn:aws:iam::961923280705:role/signUp \
--handler main --zip-file fileb://./deployment.zip

# aws lambda update-function-code --function-name signUp \
#     --zip-file fileb://./deployment.zip \
#     --region us-east-2

echo "Cleaning up"
rm -rf main 
