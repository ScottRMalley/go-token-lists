package go_token_lists

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scottrmalley/go-evm-constants/networks"
	"github.com/scottrmalley/go-token-lists/metadata"
	"math/big"
	"strings"
)

var ErrSymbolNotUnique = errors.New("could not filter for unique token with given symbol")
var ErrSymbolNotFound = errors.New("could not find token with matching symbol")
var ErrTokenNotFound = errors.New("could not find token with given address")

type List struct {
	name     metadata.Name
	metadata metadata.Metadata
	data     *ListData
}

func NewList(name metadata.Name) (*List, error) {
	md, err := metadata.TokenListMetadata(name)
	if err != nil {
		return nil, err
	}
	data, err := FetchList(context.Background(), md.Url)
	if err != nil {
		return nil, err
	}
	return &List{name, md, data}, nil
}

func (l *List) AllTokens() []Token {
	return l.data.Tokens
}

func (l *List) Tokens(network networks.Name) ([]Token, error) {
	var tokens []Token
	net, err := networks.GetNetwork(network)
	if err != nil {
		return nil, err
	}
	for _, token := range l.data.Tokens {
		if big.NewInt(token.ChainId).Cmp(net.ChainId) == 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens, nil
}

func (l *List) TokenBySymbol(network networks.Name, symbol string) (*Token, error) {
	tokens, err := l.Tokens(network)
	if err != nil {
		return nil, err
	}
	var matchingTokens []Token
	for _, token := range tokens {
		if strings.ToLower(token.Symbol) == strings.ToLower(symbol) {
			matchingTokens = append(matchingTokens, token)
		}
	}
	if len(matchingTokens) > 1 {
		return nil, ErrSymbolNotUnique
	}
	if len(matchingTokens) < 1 {
		return nil, ErrSymbolNotFound
	}
	return &matchingTokens[0], nil
}

func (l *List) TokenByAddress(network networks.Name, address common.Address) (*Token, error) {
	tokens, err := l.Tokens(network)
	if err != nil {
		return nil, err
	}
	for _, token := range tokens {
		if token.Address.Hex() == address.Hex() {
			return &token, nil
		}
	}
	return nil, ErrTokenNotFound
}

func (l *List) TokenBy(match func(*Token) bool) (*Token, error) {
	for _, token := range l.data.Tokens {
		if match(&token) {
			return &token, nil
		}
	}
	return nil, ErrTokenNotFound
}
