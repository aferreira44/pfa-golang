package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{ID: id, Price: price, Tax: tax}
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	// order.FinalPrice = order.Price + order.Tax
	return order, nil
}

func (o *Order) CalculateFinalPrice() error {
	if err := o.IsValid(); err != nil {
		return err
	}
	o.FinalPrice = o.Price + o.Tax
	return nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid ID")
	}

	if o.Price == 0 {
		return errors.New("invalid price")
	}

	if o.Tax == 0 {
		return errors.New("invalid tax")
	}
	return nil
}
