#!/bin/bash

echo "Testing the deployment url"
curl -sX GET https://7qkpsd4my6.execute-api.us-east-2.amazonaws.com/staging/cribsheets/11 | jq '.'
