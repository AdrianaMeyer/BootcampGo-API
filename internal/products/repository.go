package products

import (
	"time"
	"fmt"
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
var lastID int

type IRepository interface {
	GetAll() ([]Product, error)
	Save(id int, name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error)
	LastID() (int, error)
	Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error)
}

type repository struct {}

func NewRepository() IRepository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil

}

func (r *repository) LastID() (int, error) {
	return lastID, nil

}

func (r *repository) Save(id int, name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error) {
	
	p := Product{id, name, color, price, count, code, published, date}
	products = append(products, p)
	lastID = p.ID
	return p, nil

}

func (r *repository) Update(id int, name string, color string, price float64, count int, code string, published bool) (Product, error) {

	p := Product{Name: name, Color: color, Price: price, Count: count, Code: code, Published: published}
	updated := false

	for i := range products {
		if products[i].ID == id {
			p.ID = id
			products[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil

}