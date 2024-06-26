package repositories

import (
	"errors"
	"fmt"
	"ledger_bank/internal/domain/account"
	"ledger_bank/internal/utils"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type AccountRepository struct {
	DataBaseConection *dynamodb.DynamoDB
}

func (c *AccountRepository) CreateBankAccount(p *account.BankAccount) (*account.BankAccount, error) {
	tableName := "ledger_bank"

	av, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = c.DataBaseConection.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return p, err
}
func (c *AccountRepository) UpdateBalance(p *account.UpdateAccountParams) (*account.BankAccount, error) {
	tableName := "ledger_bank"

	acc := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ledger_bank": {
				S: aws.String(*p.Id),
			},
		},
		TableName: &tableName,
	}
	accountItem, err := c.DataBaseConection.GetItem(acc)
	if err != nil {
		return nil, err
	}
	currentBalanceAttr, ok := accountItem.Item["balance"]
	if !ok {
		return nil, errors.New("balance attribute not found")
	}
	currentVersionAttr, ok := accountItem.Item["version"]
	if !ok {
		return nil, errors.New("balance attribute not found")
	}
	currentBalance, err := strconv.ParseFloat(*currentBalanceAttr.N, 64)
	if err != nil {
		return nil, err
	}
	currentVersion, err := strconv.ParseInt(*currentVersionAttr.N, 10, 32)
	if err != nil {
		return nil, err
	}
	if currentBalance < *p.TransactionValue {
		return nil, errors.New("no enought balance")
	}
	newBalance := currentBalance + *p.TransactionValue
	newBalanceStr := strconv.FormatFloat(newBalance, 'f', -1, 64)
	currentVersionStr := strconv.FormatInt(currentVersion, 10)
	newVersion := currentVersion + 1
	newVersionStr := strconv.FormatInt(newVersion, 10)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newbal": {
				N: aws.String(newBalanceStr),
			},
			":version": {
				N: aws.String(currentVersionStr),
			},
			":newversion": {
				N: aws.String(newVersionStr),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ledger_bank": {
				S: aws.String(*p.Id),
			},
		},
		UpdateExpression:    aws.String("SET balance = :newbal, version = :newversion"),
		ConditionExpression: aws.String("version = :version"),
	}
	_, err = c.DataBaseConection.UpdateItem(input)
	if err != nil {
		return nil, err
	}
	createdDateAttr, ok := accountItem.Item["created_date"]
	if !ok {
		return nil, errors.New("balance attribute not found")
	}
	t, err := time.Parse(time.RFC3339, *createdDateAttr.S)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	var queue SendMsgQueue
	queue.Name = "ledger"
	queue.MessageBody = "Registro de transação"
	transactionValueStr := strconv.FormatFloat(*p.TransactionValue, 'f', -1, 64)
	queue.MessageGroupId = p.Id
	queue.MessageAttributes = map[string]*sqs.MessageAttributeValue{
		"id": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(*p.Id),
		},
		"value": &sqs.MessageAttributeValue{
			DataType:    aws.String("Number"),
			StringValue: aws.String(transactionValueStr),
		},
	}
	error := queue.SendMsgQueue()
	if error != nil {
		return nil, error
	}
	res := &account.BankAccount{
		Ledger_bank: *p.Id,
		Balance:     newBalance,
		CreatedDate: t,
		Version:     newVersion,
	}
	return res, nil
}
func (p *AccountRepository) UpdateListTransactions(c *account.UpdateListTransactions) error {
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
	_, err := p.DataBaseConection.UpdateItem(input)
	if err != nil {
		return err
	}
	return nil
}
func (p *AccountRepository) GetTransactions(c *account.GetTransactionParams) error
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		DataBaseConection: utils.DataBaseClient,
	}
}
