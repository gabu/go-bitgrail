package bitgrail

import (
	"context"
)

type OrderBook struct {
	Bids []OpenOrder
	Asks []OpenOrder
}

type OpenOrder struct {
	Price  float64 `json:",string"`
	Amount float64 `json:",string"`
}

// OrderBook returns bids/asks for the specified trade pair.
func (c *Client) OrderBook(ctx context.Context, pair string) (OrderBook, error) {
	req, err := c.newRequest(ctx, pair+"/orderbook", nil)
	if err != nil {
		return OrderBook{}, err
	}

	var ret = &OrderBook{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
