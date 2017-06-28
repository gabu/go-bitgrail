package bitgrail

import (
	"context"
)

type Balances map[string]Balance

type Balance struct {
	Balance  float64 `json:",string"`
	Reserved float64 `json:",string"`
}

// Balances returns balances of all coins present on BitGrail.
func (c *Client) Balances(ctx context.Context) (Balances, error) {
	req, err := c.newAuthenticatedRequest(ctx, "balances", nil)
	if err != nil {
		return Balances{}, err
	}

	var ret = &Balances{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
