package products

import (
	"time"
)

type IService interface {
	GetAll() ([]Product, error)
	Save(name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error)
	Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
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

func (s *service) Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error) {

	return s.repository.Update(id, name, color, price, count, code, published)

}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	
	return s.repository.UpdateNameAndPrice(id, name, price)

 }
 
func (s *service) Delete(id int) error {
	
	return s.repository.Delete(id)

}