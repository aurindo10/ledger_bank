package main

import (
	transactionroute "ledger_bank/internal/routes/transactionRoute"
	"ledger_bank/internal/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	utils.NewClientDb()
	lambda.Start(transactionroute.HandleCreateBankAccount)
}
