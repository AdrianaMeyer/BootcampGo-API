package tests_mocks

import (
	"errors"
)

type StubProductsSaveError struct {
}

func (s *StubProductsSaveError) Write(data interface{}) error {
	return nil
}
func (s *StubProductsSaveError) Read(data interface{}) error {
	return errors.New("JSON unexpected character")
}