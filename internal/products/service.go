package products

import "errors"

var (
	ErrNotFound = errors.New("produto não encontrado")
	NoContent = errors.New("não há produtos cadastrados")
)


type IService interface {
	GetAll() ([]Product, error)
	Save(name string, color string, price float64, count int, code string, published bool, date string) (Product, error)
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

	if len(products) == 0 {
		return nil, NoContent
	} 

	return products, nil
}

func (s *service) Save(name string, color string, price float64, count int, code string, published bool, date string) (Product, error) {

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