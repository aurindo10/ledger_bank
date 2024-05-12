package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DataBaseClient *dynamodb.DynamoDB

func NewClientDb() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	DataBaseClient = dynamodb.New(sess)
}
