# updating FindAll function 
aws lambda update-function-code --function-name FindAllCribsheets \
    --zip-file fileb://./deployment.zip \
    --region us-east-2

# Adding an Environment Variable to FindAll 
aws lambda update-function-configuration --function-name FindAllCribsheets \
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

# Adding permission to production aliased object 
aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Production" \
    --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets" \
    --principal apigateway.amazonaws.com \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414 \
    --action lambda:InvokeFunction


# Staging 
aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Staging"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets"   --principal apigateway.amazonaws.com   --statement-id e682f70d-c1a5-4042-82fc-47bb57de26f7   --action lambda:InvokeFunction
aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:${stageVariables.lambda}"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/POST/signUp"   --principal apigateway.amazonaws.com   --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414   --action lambda:InvokeFunction

aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:${stageVariables.lambda_signUp}"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/POST/signUp"   --principal apigateway.amazonaws.com   --statement-id d8612e1e-fb0a-4a9e-a8dd-663b41c721cf   --action lambda:InvokeFunction
aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:signUp:Staging"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/POST/signUp"   --principal apigateway.amazonaws.com   --statement-id 4b9c5b21-27cc-4e5d-ad2b-dea5e7e7c937   --action lambda:InvokeFunction
# Remove permission from production 
aws lambda remove-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Production" \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414

# Remove incorrect permission from Staging 
aws lambda remove-permission --function-name "arn:aws:lambda:us-east-2:961923280705:function:FindAllCribsheets:Staging" \
    --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414




########################################### FIND ONE ######################################################
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


# Updating the staging variable 
aws lambda add-permission   --function-name "arn:aws:lambda:us-east-2:961923280705:function:${stageVariables.lambda}"   --source-arn "arn:aws:execute-api:us-east-2:961923280705:7qkpsd4my6/*/GET/cribsheets"   --principal apigateway.amazonaws.com   --statement-id 80e281bc-c51d-4f53-bb96-2a95c7c26414   --action lambda:InvokeFunction



################ USEFUL CMDS ##########################
# list
aws lambda list-functions --max-items 10

# delete
aws lambda delete-function --function-name helloServerless

# Create function  
aws lambda create-function --function-name signUp --runtime go1.x \
--role arn:aws:iam::961923280705:role/signUp \
--handler main --zip-file fileb://./deployment.zip

# Add environment variable 
aws lambda update-function-configuration --function-name signUp \
 --environment Variables={COGNITO_CLIENT_ID='107fgp2jq7admcmrf19qmi7qir'} \
 --region us-east-2

# Changing the profile 
export AWS_PROFILE=whatever