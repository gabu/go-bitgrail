package bitgrail

import (
	"context"
	"fmt"
	"testing"
)

func TestOrderBook(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	ret, err := client.OrderBook(ctx, "BTC-XRB")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Bids")
	for _, bid := range ret.Bids[:2] {
		printOrder(bid)
	}
	fmt.Println("|")
	for _, bid := range ret.Bids[len(ret.Bids)-2:] {
		printOrder(bid)
	}
	fmt.Println("\nAsks")
	for _, ask := range ret.Asks[:2] {
		printOrder(ask)
	}
	fmt.Println("|")
	for _, ask := range ret.Asks[len(ret.Asks)-2:] {
		printOrder(ask)
	}
}

func printOrder(order OpenOrder) {
	fmt.Printf("Price: %.8f, Amount: %.8f\n", order.Price, order.Amount)
}
