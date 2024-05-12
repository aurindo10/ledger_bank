package routes

import (
	transactionroute "ledger_bank/internal/routes/transactionRoute"

	"github.com/aws/aws-lambda-go/lambda"
)

func AddRoutes() {
	lambda.Start(transactionroute.HandleCreateBankAccount)
}
