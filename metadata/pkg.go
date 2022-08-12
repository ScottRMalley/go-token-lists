package metadata

import (
	"errors"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var ErrListNameNotFound = errors.New("no metadata with requested name found")

type Name string

type Metadata struct {
	Networks []networks.Name
	Name     Name
	Url      string
}

func (m Metadata) Supports(network networks.Name) bool {
	for _, supported := range m.Networks {
		if supported == network {
			return true
		}
	}
	return false
}

var tokenLists = []Metadata{
	traderJoeList,
	elkFinanceList,
	pangolinList,
	pancakeswapExtendedList,
	pancakeswapDefaultList,
}

func AllLists() []Name {
	listNames := make([]Name, len(tokenLists))
	for i, meta := range tokenLists {
		listNames[i] = meta.Name
	}
	return listNames
}

func SupportedLists(network networks.Name) []Name {
	var supported []Name
	for _, list := range tokenLists {
		if list.Supports(network) {
			supported = append(supported, list.Name)
		}
	}
	return supported
}

func TokenListMetadata(name Name) (Metadata, error) {
	for _, list := range tokenLists {
		if name == list.Name {
			return list, nil
		}
	}
	return Metadata{}, ErrListNameNotFound
}
