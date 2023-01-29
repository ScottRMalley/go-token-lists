package metadata

import "github.com/scottrmalley/go-evm-constants/networks"

const UNISWAP_DEFAULT Name = "uniswap-default"

var uniswapDefaultList = Metadata{
	Networks: []networks.Name{networks.ETHEREUM_MAINNET, networks.ETHEREUM_GOERLI, networks.POLYGON_MAINNET, networks.POLYGON_TESTNET},
	Name:     UNISWAP_DEFAULT,
	Url:      "https://gateway.ipfs.io/ipns/tokens.uniswap.org",
}
