package metadata

import "github.com/scottrmalley/go-evm-constants/networks"

var TRADER_JOE_DEFAULT Name = "trader-joe-default"

var traderJoeList = Metadata{
	Networks: []networks.Name{networks.AVALANCHE_MAINNET, networks.AVALANCHE_FUJI},
	Name:     TRADER_JOE_DEFAULT,
	Url:      "https://raw.githubusercontent.com/traderjoe-xyz/joe-tokenlists/main/joe.tokenlist.json",
}

var ELK_FINANCE Name = "elk-finance"

var elkFinanceList = Metadata{
	Networks: []networks.Name{networks.AVALANCHE_MAINNET, networks.AVALANCHE_FUJI},
	Name:     ELK_FINANCE,
	Url:      "https://raw.githubusercontent.com/elkfinance/tokens/main/avax.tokenlist.json",
}

var PANGOLIN Name = "pangolin"

var pangolinList = Metadata{
	Networks: []networks.Name{networks.AVALANCHE_MAINNET, networks.AVALANCHE_FUJI},
	Name:     PANGOLIN,
	Url:      "https://raw.githubusercontent.com/pangolindex/tokenlists/main/pangolin.tokenlist.json",
}
