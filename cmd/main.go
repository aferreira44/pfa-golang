package main

import (
	"fmt"

	"github.com/aferreira44/pfa-golang/internal/order/entity"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		panic(err)
	}
	order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The final price is: %f", order.FinalPrice)
}
