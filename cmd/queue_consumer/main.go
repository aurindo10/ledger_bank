package main

import (
	"context"
	"fmt"
	"ledger_bank/internal/domain/account"
	"ledger_bank/internal/utils"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		msg := message.MessageAttributes

		fmt.Println("Mensagem:", msg)
		value := msg["value"]
		id := msg["id"]

		tableName := "ledger_bank"
		f, err := strconv.ParseFloat(*value.StringValue, 64)
		if err != nil {
			println(err.Error(), "erro 27")
		}
		request := account.UpdateListTransactions{
			Value:     &f,
			TableName: &tableName,
			Id:        *id.StringValue,
		}
		err = UpdateListTransactions(&request)
		if err != nil {
			println(err.Error(), "erro 37")
		}
	}
	return nil
}
func main() {
	utils.NewClientDb()
	lambda.Start(handler)
}
func UpdateListTransactions(c *account.UpdateListTransactions) error {
	transactionValueStr := strconv.FormatFloat(*c.Value, 'f', -1, 64)
	objects := []*dynamodb.AttributeValue{
		{
			M: map[string]*dynamodb.AttributeValue{
				"value": {
					N: aws.String(transactionValueStr),
				},
			},
		},
	}
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				L: objects,
			},
			":empty_list": {
				L: []*dynamodb.AttributeValue{},
			},
		},
		TableName: aws.String(*c.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ledger_bank": {
				S: aws.String(c.Id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set transactions = list_append(if_not_exists(transactions, :empty_list), :n)"),
	}
	_, err := utils.DataBaseClient.UpdateItem(input)
	if err != nil {
		return err
	}
	return nil
}
