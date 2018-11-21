package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// Account is the user struct for signing up
type Account struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(account Account) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}

	cognito := cognitoidentityprovider.New(cfg)
	req := cognito.SignUpRequest(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Email:    aws.String(account.Email),
		Username: aws.String(account.Username),
		Password: aws.String(account.Password),
	})
	_, err = req.Send()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(signUp)

}
