package bitgrail

import (
	"context"
)

type Ticker struct {
	Last       float64 `json:",string"`
	High       float64 `json:",string"`
	Low        float64 `json:",string"`
	Volume     float64 `json:",string"`
	CoinVolume float64 `json:",string"`
	Bid        float64 `json:",string"`
	Ask        float64 `json:",string"`
}

// Ticker returns ticker for the specified trade pair.
func (c *Client) Ticker(ctx context.Context, pair string) (Ticker, error) {
	req, err := c.newRequest(ctx, pair+"/ticker", nil)
	if err != nil {
		return Ticker{}, err
	}

	var ret = &Ticker{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
