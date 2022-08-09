# Go Token Lists

This is a Golang implementation of the [Uniswap schema](https://uniswap.org/tokenlist.schema.json) for DeFi token lists.

## Usage

```go
package main

import (
	"github.com/scottrmalley/go-evm-constants/networks"
	tl "github.com/scottrmalley/go-token-lists"
	"github.com/scottrmalley/go-token-lists/metadata"
)

func main() {
	list, err := tl.NewList(metadata.PANCAKESWAP_EXTENDED)
	if err != nil {
		panic(err)
	}

	usdc, err := list.TokenBySymbol(networks.BSC_MAINNET, "USDC")
	if err != nil {
		panic(err)
    }
	...
}
```
