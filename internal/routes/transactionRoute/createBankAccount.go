package transactionroute

import (
	"ledger_bank/internal/domain/account"
	"ledger_bank/internal/repositories"
	"ledger_bank/internal/utils"

	"github.com/aws/aws-lambda-go/events"
)

func HandleCreateBankAccount(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	result, problems, error := utils.DecodeValid[account.BankAccountParams](&request)
	if error != nil {
		if len(problems) > 0 {
			res := utils.Encode(400, problems)
			return res, nil
		}
		res := utils.Encode(400, problems)
		return res, nil
	}
	repo := repositories.NewAccountRepository()
	domain := account.NewCreateAccountDomain(repo)
	r, err := domain.Execute(&result)
	if err != nil {
		res := utils.Encode(400, "Erro ao criar registro")
		return res, nil
	}
	res := utils.Encode(201, r)
	return res, nil
}
