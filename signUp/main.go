package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// Account is the user struct for signing up
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(account Account) (events.APIGatewayProxyResponse, error) {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, err
	}

	cognito := cognitoidentityprovider.New(cfg)
	req := cognito.SignUpRequest(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username: aws.String(account.Username),
		Password: aws.String(account.Password),
	})
	_, err = req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while creating new user in Cognito Pool",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(signUp)

}
