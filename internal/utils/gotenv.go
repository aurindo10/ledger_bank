package utils

import (
	"os"

	"github.com/gofor-little/env"
)

func NewEnv(key string) (*string, error) {
	error := env.Load("../../.env")

	if error != nil {
		println("Erro ao carregar as variáveis de ambiente")
		return nil, error
	}
	env := os.Getenv(key)
	return &env, nil
}
