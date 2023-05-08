package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"store/api"
	"store/internal/service/product/model"

	"github.com/gorilla/mux"
)

type resource struct {
	m map[int]*model.Product
}

// list products
// create product
// read product
// update product

func RegisterHandlers(r *mux.Router, m map[int]*model.Product) {

	res := resource{m: m}

	r.HandleFunc("/products", res.GetProducts).Methods("GET")
	r.HandleFunc("/products", res.AddProduct).Methods("POST")
	r.HandleFunc("/products/{id}", res.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", res.GetProduct).Methods("GET")

}

func (res resource) AddProduct(w http.ResponseWriter, r *http.Request) {
	req := &model.Product{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("error in decoding", err)
	}
	p := model.ProductReq{
		Product: *req,
	}
	res.m[p.Product.ProductId] = &p.Product
	api.JsonResponse(w, http.StatusOK, api.NewResponse(true, "product is added successfylly", nil))
}

func (res resource) GetProducts(w http.ResponseWriter, r *http.Request) {

	products := model.ProductRes{}

	for _, p := range res.m {
		products.Products = append(products.Products, *p)

	}
	api.JsonResponse(w, http.StatusOK, products)
}

func (res resource) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	prodId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invlid input")
	}

	for _, p := range res.m {
		if p.ProductId == prodId {
			api.JsonResponse(w, http.StatusOK, p)
			return
		}

	}
	api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "product is not available", fmt.Errorf("product is not available")))
}

func (res resource) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	req := &model.Product{}
	id := mux.Vars(r)["id"]
	prodId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invlid input")
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("error in decoding", err)
	}

	for _, p := range res.m {
		if p.ProductId == prodId {
			p := model.ProductReq{
				Product: model.Product{
					ProductId:    prodId,
					ProductName:  req.ProductName,
					Quantity:     req.Quantity,
					Avaliability: req.Avaliability,
					Price:        req.Price,
					Category:     req.Category,
				},
			}
			res.m[prodId] = &p.Product
			api.JsonResponse(w, http.StatusOK, api.NewResponse(true, "product is updated successfylly", nil))
			return
		}

	}
	api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "product is not available", fmt.Errorf("product is not available")))
}
