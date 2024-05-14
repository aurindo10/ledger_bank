package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var DataBaseClient *dynamodb.DynamoDB
var SqsConnection *sqs.SQS

func NewClientDb() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	SqsConnection = sqs.New(sess)
	DataBaseClient = dynamodb.New(sess)
}
