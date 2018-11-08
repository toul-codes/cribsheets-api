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
	ID           string `json:"id"`
	Professor    string `json:"professor"`
	CourseNumber string `json:"coursenumber"`
	Description  string `json:"description"`
	Filename     string `json:"filename"`
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
			"Professor": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Professor),
			},
			"CourseNumber": dynamodb.AttributeValue{
				S: aws.String(cribsheet.CourseNumber),
			},
			"Description": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Description),
			},
			"Filename": dynamodb.AttributeValue{
				S: aws.String(cribsheet.Filename),
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
