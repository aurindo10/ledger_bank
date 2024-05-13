package account

import (
	"reflect"
	"time"
)

type BankAccount struct {
	Ledger_bank string    `json:"ledger_bank"`
	Balance     float64   `json:"balance"`
	CreatedDate time.Time `json:"created_date"`
	Version     int64     `json:"version"`
}

type BankAccountParams struct {
	Balance     *float64   `json:"balance"`
	CreatedDate *time.Time `json:"created_date"`
}

func (p BankAccountParams) Valid() (problems map[string]string) {
	problems = make(map[string]string)

	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	// Verifica se o dado passado é uma estrutura
	if v.Kind() != reflect.Struct {
		problems["error"] = "O tipo de dados não é uma estrutura"
		return problems
	}

	// Itera sobre os campos da estrutura
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Verifica se o campo é nulo
		if field.IsNil() {
			problems[fieldName] = "Campo está nulo"
		}
	}

	return problems
}

type UpdateAccountParams struct {
	TransactionValue *float64 `json:"transaction_value"`
	Id               *string  `json:"id"`
}

func (p UpdateAccountParams) Valid() (problems map[string]string) {
	problems = make(map[string]string)

	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	// Verifica se o dado passado é uma estrutura
	if v.Kind() != reflect.Struct {
		problems["error"] = "O tipo de dados não é uma estrutura"
		return problems
	}

	// Itera sobre os campos da estrutura
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Verifica se o campo é nulo
		if field.IsNil() {
			problems[fieldName] = "Campo está nulo"
		}
	}

	return problems
}

type UpdateListTransactions struct {
	Value     *float64 `json:"value"`
	TableName *string
	Id        string
}
type Repository interface {
	CreateBankAccount(p *BankAccount) (*BankAccount, error)
	UpdateBalance(p *UpdateAccountParams) (*BankAccount, error)
}
