package transactionroute

import (
	"context"
	"fmt"
	"ledger_bank/internal/domain/account"
	"ledger_bank/internal/repositories"
	"ledger_bank/internal/utils"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func HandlerConsumeSqs(ctx context.Context, sqsEvent events.SQSEvent) error {
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
		repo := repositories.NewAccountRepository()
		domain := account.NewUpdateTransaction(repo)
		err = domain.Execute(&request)
		if err != nil {
			return err
		}
		queue := "ledger"
		result, err := utils.SqsConnection.GetQueueUrl(&sqs.GetQueueUrlInput{
			QueueName: &queue,
		})
		if err != nil {
			println(err.Error(), "erro 37")
		}
		receiptHandle := message.ReceiptHandle
		utils.SqsConnection.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      result.QueueUrl,
			ReceiptHandle: &receiptHandle,
		})
		if err != nil {
			println(err.Error(), "erro 50")
		}
		println("item da fila deletado com sucesso!")
	}
	return nil
}
