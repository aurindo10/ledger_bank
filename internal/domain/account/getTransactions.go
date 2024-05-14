package account

type GetTransactions struct {
	repository Repository
}

func (c *GetTransactions) Execute(p *GetTransactionParams) error {
	err := c.repository.GetTransactions(p)
	if err != nil {
		return err
	}
	return nil
}

func NewGetTransactions(repository Repository) *GetTransactions {
	return &GetTransactions{
		repository: repository,
	}
}
