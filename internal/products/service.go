package products

import (
	"time"
)

type IService interface {
	GetAll() ([]Product, error)
	Save(name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error)
}

type service struct {
	repository IRepository
}

func NewService(r IRepository) IService {
	return &service {
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Save(name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error) {

	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Save(lastID, name, color, price, count, code, published, date)
	if err != nil {
		return Product{}, err
	}

	return product, nil

}
