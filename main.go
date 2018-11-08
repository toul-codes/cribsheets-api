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

// Cribsheet struct is used for storing the data
// and marshalling it into JSON for the request body
type Cribsheet struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Subject     string `json:"subject"`
	Type        string `json:"type"`
	University  string `json:"university"`
	Year        string `json:"year"`
	Description string `json:"description"`
}

// findAll will scan the DB and return all of the items in the DB
func findAll(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})

	res, err := req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		}, nil
	}

	cribsheets := make([]Cribsheet, 0)
	for _, item := range res.Items {
		cribsheets = append(cribsheets, Cribsheet{
			ID:          *item["ID"].S,
			Name:        *item["Name"].S,
			LastName:    *item["LastName"].S,
			Subject:     *item["Subject"].S,
			Type:        *item["Type"].S,
			University:  *item["University"].S,
			Year:        *item["Year"].S,
			Description: *item["Description"].S,
		})
	}

	response, err := json.Marshal(cribsheets)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findAll)
}
