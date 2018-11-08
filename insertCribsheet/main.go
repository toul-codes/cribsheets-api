package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Cribsheet in a data holding entitity
type Cribsheet struct {
	ID         string `json:"id"`
	LastName   string `json:"lastname"`
	Subject    string `json:"subject"`
	Type       string `json:"type"`
	University string `json:"university"`
	Year       string `json:"year"`
}

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var cribsheet Cribsheet

	err := json.Unmarshal([]byte(request.Body), &cribsheet)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid payload",
		}, nil
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)

	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(cribsheet.ID),
			},
			"LastName": dynamodb.AttributeValue{
				S: aws.String(cribsheet.LastName),
			},
			"Subject": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Subject),
			},
			"Type": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Type),
			},
			"University": dynamodb.AttributeValue{
				S: aws.String(cribsheet.University),
			},
			"Year": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Year),
			},
		},
	})
	_, err = req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while inserting Cribsheet to DynamoDB",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(insert)
}
