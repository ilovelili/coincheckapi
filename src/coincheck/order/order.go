package order

import (
	"coincheck"

	"encoding/json"
	"errors"
)

// NewOrder sends a new order.
func NewOrder(api *coincheck.APIClient, order Order) (newOrder Order, err error) {
	newOrder = order
	data, err := json.Marshal(newOrder)
	if err != nil {
		return newOrder, err
	}
	err = api.DoPostRequest(NewOrderEndpoint, data, &newOrder)
	if err != nil {
		return newOrder, err
	}
	if newOrder.Error != "" {
		return newOrder, errors.New(newOrder.Error)
	}
	return newOrder, nil
}
