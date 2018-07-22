package chest

import (
	"strconv"
)

type CollectionService struct {
	client *Client
}

type CollectionCard struct {
	Id          int    `json:"card_id"`
	Name        string `json:"card_name"`
	EditionCode string `json:"edition_code"`
	Count       int    `json:"card_count"`
}

type Collection struct {
	Cards      []*CollectionCard `json:"items"`
	Pagination *Pagination       `json:"pagination"`
}

func (s *CollectionService) UserCollection(params TableParams) (
	*Collection, *Response, error,
) {
	req, err := s.client.NewRequest("GET", "/api/v1/collection", nil)

	if err != nil {
		return nil, nil, err
	}

	if params.Page > 1 {
		query := req.URL.Query()
		query.Add("page", strconv.Itoa(params.Page))
		req.URL.RawQuery = query.Encode()
	}

	f := &Collection{}
	resp, err := s.client.Do(req, f)

	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
