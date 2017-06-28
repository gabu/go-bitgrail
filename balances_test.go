package bitgrail

import (
	"context"
	"fmt"
	"testing"
)

func TestBalances(t *testing.T) {
	client := newAuthClient()
	ctx := context.Background()
	ret, err := client.Balances(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
