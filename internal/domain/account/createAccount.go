package account

import "github.com/google/uuid"

type CreateAccountDomain struct {
	repository Repository
}

func (r CreateAccountDomain) Execute(p *BankAccountParams) (*BankAccount, error) {
	id := uuid.New().String()
	account := &BankAccount{
		Ledger_bank: id,
		Balance:     *p.Balance,
		CreatedDate: *p.CreatedDate,
		Version:     1,
	}
	result, err := r.repository.CreateBankAccount(account)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCreateAccountDomain(repository Repository) *CreateAccountDomain {
	return &CreateAccountDomain{
		repository: repository,
	}
}
