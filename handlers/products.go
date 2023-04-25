package handlers

import (
	"log"
	"net/http"

	"github.com/ellofae/RESTful-API-for-Bakery-microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET method")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Didn't manage to encode products data", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST method")

	productObj := &data.Product{}

	err := productObj.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Didn't manage to decode product's data", http.StatusInternalServerError)
		return
	}

	data.AddProduct(productObj)
}
