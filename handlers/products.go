package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT method")

		reg := regexp.MustCompile("/([0-9]+)")
		d := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(d) != 1 {
			p.l.Println("Incorrect URI: didn't manage to recive any data")
			http.Error(rw, "Incorrect URI", http.StatusBadRequest)
			return
		}

		if len(d[0]) != 2 {
			p.l.Println("Incorrect URI: no data was recived from the r.URL.Path")
			http.Error(rw, "Incorrect URI", http.StatusBadRequest)
			return
		}

		uriID := d[0][1]
		id, err := strconv.Atoi(uriID)
		if err != nil {
			p.l.Println("Incorrect URI: not possible to convert recived id from string to int")
			http.Error(rw, "Incorrect URI", http.StatusBadRequest)
		}
		p.updateData(id, rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusNotImplemented)
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

func (p *Products) updateData(id int, rw http.ResponseWriter, r *http.Request) {
	productObj := &data.Product{}

	err := productObj.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Didn't manage to decode data", http.StatusInternalServerError)
		return
	}

	err = data.UpdateData(id, productObj)
	if err == data.ErrProductNotFound {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}
}
