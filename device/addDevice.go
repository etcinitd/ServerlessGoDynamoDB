package main

import (
	"context"
	"fmt"
	"os"
	"encoding/json"
        "strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Device struct {
	ID          string      `json:"id"`
	DeviceModel string 	`json:"devicemodel"`
	Name        string   	`json:"name"`
	Note        string 	`json:"note"`
	Serial      string 	`json:"serial"`
}

var ddb *dynamodb.DynamoDB
func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ 
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) 
	}
}

func AddDevice(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var (
                device Device
		tableName = aws.String(os.Getenv("DEVICE_TABLE_NAME"))
	)

	// Parse request body
	json.Unmarshal([]byte(request.Body),&device)

        if empty(device.ID) {
                return events.APIGatewayProxyResponse{
                        Body: "Bad Request: Id Can Not Be Empty", 
                        StatusCode: 400,
                }, nil
	}

        if empty(device.DeviceModel) {
                return events.APIGatewayProxyResponse{
                        Body: "Bad Request: DeviceModel Can Not Be Empty",
                        StatusCode: 400,
                }, nil
        }

        if empty(device.Name) {
                return events.APIGatewayProxyResponse{
                        Body: "Bad Request: Name Can Not Be Empty",
                        StatusCode: 400,
                }, nil
        }

        if empty(device.Note) {
                return events.APIGatewayProxyResponse{
                        Body: "Bad Request: Note Can Not Be Empty",
                        StatusCode: 400,
                }, nil
        }

        if empty(device.Serial) {
                return events.APIGatewayProxyResponse{
                        Body: "Bad Request: Serial Can Not Be Empty",
                        StatusCode: 400,
                }, nil
        }

	item, err := dynamodbattribute.MarshalMap(device)
        if err != nil {
                return events.APIGatewayProxyResponse{ 
                        Body: "Internal Server Error",
                        StatusCode: 500,
                }, nil
        }

	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: tableName,
	}

	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ 
			Body: "Internal Server Error",
			StatusCode: 500,
		}, nil
	} 

	body, err := json.Marshal(device)
        if err != nil {
       		return events.APIGatewayProxyResponse{ 
                       	Body: "Internal Server Error",
                       	StatusCode: 500,
               	}, nil
        }

	return events.APIGatewayProxyResponse{ // Success HTTP response
		Body: string(body),
		StatusCode: 201,
	}, nil
}

func empty(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}

func main() {
	lambda.Start(AddDevice)
}
