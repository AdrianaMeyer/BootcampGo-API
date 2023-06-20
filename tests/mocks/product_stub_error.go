package tests_mocks

import (
	"errors"
)

type StubProductsError struct {
}

func (s *StubProductsError) Write(data interface{}) error {
	return nil
}
func (s *StubProductsError) Read(data interface{}) error {
	return errors.New("Erro ao ler os dados")
}