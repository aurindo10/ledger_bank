package account

type UpdateAccount struct {
	repository Repository
}

func (p UpdateAccount) Execute(r *UpdateAccountParams) (*BankAccount, error) {
	rep, error := p.repository.UpdateBalance(r)
	if error != nil {
		return nil, error
	}
	return rep, nil

}

func NewUpdateAccount(repository Repository) *UpdateAccount {
	return &UpdateAccount{
		repository: repository,
	}
}
