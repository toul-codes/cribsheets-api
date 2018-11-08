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

// Cribsheet app represents the structure that stores the material's metadata
type Cribsheet struct {
	ID           string `json:"id"`
	Professor    string `json:"professor"`
	CourseNumber string `json:"coursenumber"`
	Description  string `json:"description"`
	Filename     string `json:"filename"`
}

func findOne(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	})
	res, err := req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while fetching Cribsheet from DynamoDB",
		}, nil
	}

	cribsheet := Cribsheet{
		ID:           *res.Item["ID"].S,
		Professor:    *res.Item["Professor"].S,
		CourseNumber: *res.Item["CourseNumber"].S,
		Description:  *res.Item["Description"].S,
		Filename:     *res.Item["Filename"].S,
	}

	response, err := json.Marshal(cribsheet)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findOne)
}
