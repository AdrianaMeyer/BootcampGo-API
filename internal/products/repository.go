package products

import (
	"time"
	"fmt"

	"github.com/AdrianaMeyer/BootcampGo-API/pkg/store"
)

type Product struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Color 		string 		`json:"color"`
	Price 		float64 	`json:"price"`
	Count 		int 		`json:"count"`
	Code 		string 		`json:"code"`
	Published 	bool 		`json:"published"`
	Date 		time.Time 	`json:"date"`
}

var products []Product

type IRepository interface {
	GetAll() ([]Product, error)
	Save(id int, name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error)
	LastID() (int, error)
	Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product

   	err := r.db.Read(&products)
	if err != nil {
		return nil, err
	 }

   	return products, nil

}

func (r *repository) LastID() (int, error) {
	var products []Product

	err := r.db.Read(&products); 
	if err != nil {
       return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	lastProduct := products[len(products)-1]
	return lastProduct.ID, nil

}

func (r *repository) Save(id int, name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error) {
	
	var products []Product
	r.db.Read(&products)

	p := Product{id, name, color, price, count, code, published, date}
	products = append(products, p)

	err := r.db.Write(products)
	if err != nil {
		return Product{}, err
	}

	return p, nil

}

func (r *repository) Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error) {

	p := Product{Name: name, Color: color, Price: price, Count: count, Code: code, Published: published}
	updated := false

	for i := range products {
		if products[i].ID == id {
			p.ID = id
			p.Date = products[i].Date
			products[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil

}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	var p Product
	updated := false
	for i := range products {
		if products[i].ID == id {
			products[i].Name = name
			products[i].Price = price
			updated = true
			p = products[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d no encontrado", id)
	}
	return p, nil
 }
 

 func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range products {
		if products[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Produto %d nao encontrado", id)
	}
	
	products = append(products[:index], products[index+1:]...)
	return nil
 }
 