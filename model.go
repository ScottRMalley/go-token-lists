package go_token_lists

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type Token struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Address  common.Address `json:"address"`
	ChainId  int64          `json:"chainId"`
	Decimals uint           `json:"decimals"`
	LogoURI  string         `json:"logoURI"`
}

type ListData struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Version   Version   `json:"version"`
	LogoURI   string    `json:"logoURI"`
	Keywords  []string  `json:"keywords"`
	Tokens    []Token   `json:"tokens"`
}
