package metadata

import "github.com/scottrmalley/go-evm-constants/networks"

const PANCAKESWAP_EXTENDED Name = "pancakeswap-extended"

var pancakeswapExtendedList = Metadata{
	Networks: []networks.Name{networks.BSC_MAINNET},
	Name:     PANCAKESWAP_EXTENDED,
	Url:      "https://tokens.pancakeswap.finance/pancakeswap-extended.json",
}

const PANCAKESWAP_DEFAULT Name = "pancakeswap-default"

var pancakeswapDefaultList = Metadata{
	Networks: []networks.Name{networks.BSC_MAINNET},
	Name:     PANCAKESWAP_DEFAULT,
	Url:      "https://tokens.pancakeswap.finance/pancakeswap-default.json",
}
