package order

import (
	"fmt"
	"store/internal/service/order/repo"
)

func Get() {

	fmt.Println("Get method called for order servie")
	repo.Get()

}
