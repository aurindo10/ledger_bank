package transactionroute

import (
	"ledger_bank/internal/domain/account"
	"ledger_bank/internal/repositories"
	"ledger_bank/internal/utils"

	"github.com/aws/aws-lambda-go/events"
)

func HandleCreateTransaction(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	decoded, problems, error := utils.DecodeValid[account.UpdateAccountParams](&request)
	if error != nil {
		if len(problems) > 0 {
			res := utils.Encode(400, problems)
			return res, nil
		}
		res := utils.Encode(400, problems)
		return res, nil
	}
	repo := repositories.NewAccountRepository()
	domain := account.NewUpdateAccount(repo)
	result, err := domain.Execute(&decoded)
	if err != nil {
		res := utils.Encode(500, err.Error())
		return res, nil
	}
	res := utils.Encode(200, result)
	return res, nil
}
