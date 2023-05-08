package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"store/api"
	"store/internal/service/order/model"
	pmodel "store/internal/service/product/model"
	"strconv"

	"github.com/gorilla/mux"
)

type resource struct {
	m map[int]*pmodel.Product
	o map[int]*model.OrderRes
}

var orderId int

func NewOrder() int {
	orderId = orderId + 1
	return orderId
}

func RegisterHandlers(r *mux.Router, m map[int]*pmodel.Product, o map[int]*model.OrderRes) {

	res := resource{m: m, o: o}

	r.HandleFunc("/order", res.PlaceOrder).Methods("POST")
	r.HandleFunc("/order", res.GetOrders).Methods("GET")
	r.HandleFunc("/order/{id}", res.GetOrder).Methods("GET")
	r.HandleFunc("/order/{id}", res.updateOrder).Methods("PATCH")
}

func (res resource) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	req := model.OrderReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("error in decoding", err)
	}

	resp := model.OrderResponse{}
	resp.Total = 0.0
	for _, o := range req.Orders {
		v, ok := res.m[o.ProductId]
		if !ok {
			api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "product is not available", fmt.Errorf("product is not available")))
			return
		}
		if v.Quantity < o.Quantity {
			api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "required quantity is not available", fmt.Errorf("required quantity is not available")))
			return
		}
		if o.Quantity > 10 {
			api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "max quantity allowed is 10", fmt.Errorf("max quantity allowed is 10")))
			return
		}
		prodId := o.ProductId
		prod, ok := res.m[prodId]
		if ok {
			prod.Quantity -= o.Quantity
		}
		orderId := NewOrder()
		resp.Total = resp.Total + v.Price*float64(o.Quantity)
		resp.DiscountPrice = 0.0
		resp.Success = true
		order := model.OrderRes{
			OrderId:     orderId,
			ProductId:   v.ProductId,
			ProductName: v.ProductName,
			Quantity:    o.Quantity,
			Price:       v.Price,
			OrderStatus: model.Placed,
		}
		resp.Orders = append(resp.Orders, order)
		res.o[orderId] = &order
	}

	if len(req.Orders) >= 3 {
		resp.Total = resp.Total - (resp.Total / 10)
		resp.DiscountPrice = resp.Total / 10
	}

	api.JsonResponse(w, http.StatusOK, resp)

}

func (res resource) GetOrders(w http.ResponseWriter, r *http.Request) {

	orders := model.Orders{}

	for _, o := range res.o {
		orders.Orders = append(orders.Orders, *o)

	}
	api.JsonResponse(w, http.StatusOK, orders)
}

func (res resource) GetOrder(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	orderId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invlid input")
	}
	for _, o := range res.o {
		if o.OrderId == orderId {
			api.JsonResponse(w, http.StatusOK, o)
			return
		}
	}
	api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "order id is not available", fmt.Errorf("order id is not available")))
}

func (res resource) updateOrder(w http.ResponseWriter, r *http.Request) {

	req := &model.OrderStatusReq{}
	id := mux.Vars(r)["id"]

	orderId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invlid input")
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("error in decoding", err)
	}

	for _, o := range res.o {
		if o.OrderId == orderId {
			order := model.OrderRes{
				OrderId:     orderId,
				ProductId:   o.ProductId,
				ProductName: o.ProductName,
				Quantity:    o.Quantity,
				Price:       o.Price,
				OrderStatus: req.OrderStatus,
			}
			res.o[orderId] = &order
			api.JsonResponse(w, http.StatusOK, api.NewResponse(true, "order is updated successfylly", nil))
			return
		}
	}
}
