package go_token_lists

import (
	"context"
	"github.com/carlmjohnson/requests"
)

func FetchList(ctx context.Context, url string) (*ListData, error) {
	var response ListData
	err := requests.
		URL(url).
		ToJSON(&response).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
