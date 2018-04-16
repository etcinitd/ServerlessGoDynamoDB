package main

import (
	"context"
	"fmt"
	"encoding/json"
	"os"

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
	Name        string  	`json:"name"`
	Note        string 	`json:"note"`
        Serial      string      `json:"serial"`
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

func GetDevice(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var (
		tableName = aws.String(os.Getenv("DEVICE_TABLE_NAME"))
                id = request.PathParameters["id"]
	)

        input := &dynamodb.GetItemInput{
       		Key: map[string]*dynamodb.AttributeValue{
             		"id": {
                		S: aws.String(id),
                	},
                },
                TableName: tableName, 
         }

         result, err := ddb.GetItem(input)

         if err != nil {
                return events.APIGatewayProxyResponse{
                        Body: "Internal Server Error",
                        StatusCode: 500,
                        }, nil
         } 

	 if len(result.Item) <= 0  {
         	return events.APIGatewayProxyResponse{ 
                	Body: "Not Found",
                	StatusCode: 404,
                	}, nil 
         }

         device := Device{}
         if err := dynamodbattribute.UnmarshalMap(result.Item, &device); err != nil {
                return events.APIGatewayProxyResponse{
                        Body: "Internal Server Error",
                        StatusCode: 500,
                        }, nil
         }

         body, err := json.Marshal(device);
         if err != nil {
                return events.APIGatewayProxyResponse{
                        Body: "Internal Server Error",
                        StatusCode: 500,
                        }, nil
         }

	 return events.APIGatewayProxyResponse{ // Success HTTP response
		Body: string(body),
		StatusCode: 200,
         	}, nil
}

func main() {
	lambda.Start(GetDevice)
}
