package domain

type Request struct {
	Name 		string 		`json:"name"`
	Color 		string 		`json:"color"`
	Price 		float64 	`json:"price"`
	Count 		int 		`json:"count"`
	Code 		string 		`json:"code"`
	Published 	bool 		`json:"published"`
}

type RequestUpdateNameAndPrice struct {
	Name 		string 		`json:"name"`
	Price 		float64 	`json:"price"`
}
