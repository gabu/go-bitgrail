# go-bitgrail

An unofficial [BitGrail Public and Private API](https://bitgrail.com/api-documentation) client for Go.

## Supports

### Public API

- [x] Ticker
- [x] Order book
- [ ] Trade history

### Private API (needs an authentication)

- [x] Balances
- [ ] Buy order
- [ ] Sell order
- [ ] Open orders
- [ ] Cancel order
- [ ] Get deposit address
- [ ] Withdraw
- [ ] Last trades
- [ ] Deposits history
- [ ] Withdraws history


## Usage

### Ticker

```go
package main

import (
	"context"
	"fmt"

	"github.com/gabu/go-bitgrail"
)

func main() {
	client := bitgrail.NewClient()
	ctx := context.Background()
	ticker, err := client.Ticker(ctx, "BTC-XRB")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Last: %.8f, High: %.8f, Low: %.8f, Volume: %.8f, CoinVolume: %.8f, Bid: %.8f, Ask: %.8f\n",
		ticker.Last, ticker.High, ticker.Low, ticker.Volume, ticker.CoinVolume, ticker.Bid, ticker.Ask)
}
```

### Authentication

```go
func main() {
	client := bitgrail.NewClient().Auth("YOUR API KEY", "YOUR API SECRET")
}
```
