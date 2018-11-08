#!/bin/bash
curl -sX POST -d '{"id":"17","name":"Exam 1 Review", "lastname":"Czuron", "subject":"MATH", "type":"hw", "university":"University of Poland", "year":"1999", "description":"has solutions"}' https://7qkpsd4my6.execute-api.us-east-2.amazonaws.com/staging/cribsheets | jq '.'
