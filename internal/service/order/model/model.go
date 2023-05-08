package model

type OrderStatus string

const (
	Placed     OrderStatus = "Placed"
	Dispatched OrderStatus = "Dispatched"
	Completed  OrderStatus = "Completed"
	Cancelled  OrderStatus = "Cancelled"
)

type Order struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type OrderReq struct {
	Orders []Order
}

type OrderResponse struct {
	Success       bool    `json:"success"`
	Total         float64 `json:"total"`
	DiscountPrice float64 `json:"discountPrice"`
	Orders        []OrderRes
}

type OrderRes struct {
	OrderId     int         `json:"orderId"`
	ProductId   int         `json:"productId"`
	ProductName string      `json:"productName"`
	Quantity    int         `json:"quantity"`
	Price       float64     `json:"price"`
	OrderStatus OrderStatus `json:"orderStatus"`
}

type Orders struct {
	Orders []OrderRes `json:"orders"`
}

type OrderStatusReq struct {
	OrderStatus OrderStatus `json:"orderStatus"`
}
