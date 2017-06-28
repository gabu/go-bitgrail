package bitgrail

import (
	"context"
	"fmt"
	"testing"
)

func TestTicker(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	ret, err := client.Ticker(ctx, "BTC-XRB")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Last: %.8f, High: %.8f, Low: %.8f, Volume: %.8f, CoinVolume: %.8f, Bid: %.8f, Ask: %.8f\n",
		ret.Last, ret.High, ret.Low, ret.Volume, ret.CoinVolume, ret.Bid, ret.Ask)
}
