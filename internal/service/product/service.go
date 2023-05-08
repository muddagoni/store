package product

import (
	"fmt"
	"store/internal/service/product/repo"
)

func Get() {

	fmt.Println("Get method called for product servie")
	repo.Get()

}
