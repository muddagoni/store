package model

type Category string

const (
	Premium Category = "Premium"
	Regular Category = "Regular"
	Budget  Category = "Budget"
)

type Product struct {
	ProductId    int      `json:"productId"`
	ProductName  string   `json:"productName"`
	Quantity     int      `json:"quantity"`
	Avaliability bool     `json:"availability"`
	Price        float64  `json:"price"`
	Category     Category `json:"category"`
}

type ProductReq struct {
	Product Product
}

type ProductRes struct {
	Products []Product `json:"products"`
}
