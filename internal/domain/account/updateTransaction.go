package account

type UpdateTransaction struct {
	repository Repository
}

func (c *UpdateTransaction) Execute(p *UpdateListTransactions) error {
	err := c.repository.UpdateListTransactions(p)
	if err != nil {
		return err
	}
	return nil
}
func NewUpdateTransaction(repository Repository) *UpdateTransaction {
	return &UpdateTransaction{
		repository: repository,
	}
}
