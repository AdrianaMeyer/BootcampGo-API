package products

import (
	"time"
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

type Repository interface {
	GetAll() ([]Product, error)
	Save(id int, name string, color string, price float64, count int, code string, published bool, date time.Time) (Product, error)
	LastID() (int, error)
}

type repository struct {}

func NewRepository() Repository {
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
