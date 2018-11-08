#!/bin/bash

echo "Testing the staging environment url"
curl -sX GET https://7qkpsd4my6.execute-api.us-east-2.amazonaws.com/staging/cribsheets/ | jq '.'

echo "Testing the production  environment url"
curl -sX GET https://7qkpsd4my6.execute-api.us-east-2.amazonaws.com/production/cribsheets/ | jq '.'