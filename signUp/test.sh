#!/bin/bash
curl -sX POST -d '{"id":"15","Professor":"Rhoden", "coursenumber":"3339", "description":"Exam 1", "filename":"exam 1 review"}' https://7qkpsd4my6.execute-api.us-east-2.amazonaws.com/staging/cribsheets | jq '.'

# useless change for a commit to trigger a build 
cd code

