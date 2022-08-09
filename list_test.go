package go_token_lists

import (
	"github.com/scottrmalley/go-token-lists/metadata"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ListTestSuite struct {
	suite.Suite
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (s *ListTestSuite) TestListFetching() {
	t := s.T()
	listNames := []metadata.Name{
		metadata.TRADER_JOE_DEFAULT,
		metadata.PANCAKESWAP_DEFAULT,
		metadata.PANCAKESWAP_EXTENDED,
		metadata.PANGOLIN,
		metadata.ELK_FINANCE,
	}
	t.Run("should be able to fetch a metadata from every preconfigured metadata name", func(t *testing.T) {
		for _, listName := range listNames {
			list, err := NewList(listName)
			require.NoErrorf(t, err, "could not fetch list for name: %s", listName)
			require.NotEmpty(t, list, "returned list was empty")
		}
	})
}
