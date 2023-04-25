package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var productList = []*Product{
	&Product{
		ID:          1,
		Title:       "Chocolate cake",
		Description: "A fluffy cake made of Alpine dark chocolate",
		Price:       5.99,
		SKU:         "soag214f",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Title:       "Brownie",
		Description: "A tasty chocolate make with berries",
		Price:       1.99,
		SKU:         "fas412a",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Title:       "Croissant",
		Description: "A crunchy delisious make of bread and vanilla",
		Price:       2.99,
		SKU:         "opf123h",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
