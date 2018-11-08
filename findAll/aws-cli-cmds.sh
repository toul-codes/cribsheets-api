# updating FindAll function 
aws lambda update-function-code --function-name FindAllCribsheets \
    --zip-file fileb://./deployment.zip \
    --region us-east-2

# Adding an Environment Variable 
aws lambda update-function-configuration --function-name FindAllCribsheets \
    --environment Variables={TABLE_NAME=cribsheets} \
    --region us-east-2

# Update FindOneCribsheet 
aws lambda update-function-code --function-name FindOneCribsheet \
    --zip-file fileb://./deployment.zip \
    --region us-east-2

# Update FindOneCribsheet with Environment table name
aws lambda update-function-configuration --function-name FindOneCribsheet \
 --environment Variables={TABLE_NAME=cribsheets} \
 --region us-east-2

# Update InsertOneCribsheet with env. table name
aws lambda update-function-configuration --function-name InsertCribsheet \
 --environment Variables={TABLE_NAME=cribsheets} \
 --region us-east-2

# versioning 
aws lambda publish-version --function-name FindAllCribsheets --description 1.1.0

# Calling and older version 
aws lambda invoke --function-name FindAllCribsheets --qualifier 1 result.json

# Testing
aws lambda invoke --function-name FindAllCribsheets --payload file://input.json --qualifier 2 result.json

# Aliasing for Versioning Pointers 
aws lambda create-alias --function-name FindAllCribsheets \
    --name Production --description "Production environment" \
    --function-version 1

# Updating the staging variable 
aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:${stageVariables.lambda}"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets"   --principal apigateway.amazonaws.com   --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414   --action lambda:InvokeFunction

# Production Alias 
aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Production" \
    --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets" \
    --principal apigateway.amazonaws.com \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414 \
    --action lambda:InvokeFunction

# Staging 
aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Staging" \
    --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets" \
    --principal apigateway.amazonaws.com \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414 \
    --action lambda:InvokeFunction

aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Staging"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets"   --principal apigateway.amazonaws.com   --statement-id fb583b09-c1fc-4c80-a3ed-4381978cc2df   --action lambda:InvokeFunction

# Remove permission from production 
aws lambda remove-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Production" \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414

# Remove incorrect permission from Staging 
aws lambda remove-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Staging" \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414
