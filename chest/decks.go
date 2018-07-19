package chest

import (
	"fmt"
)

type DeckService struct {
	client *Client
}

type Deck struct {
	Id   int    `json:"deck_id"`
	Name string `json:"deck_name"`

	Cards []*Card `json:"cards"`
}

type DeckList struct {
	Decks      []*Deck    `json:"items"`
	Pagination Pagination `json:"pagination"`
}

func (s *DeckService) UserDecks() (*DeckList, *Response, error) {
	req, err := s.client.NewRequest("GET", "/api/v1/decks", nil)

	if err != nil {
		return nil, nil, err
	}

	f := &DeckList{}
	resp, err := s.client.Do(req, f)

	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}

func (s *DeckService) UserDeck(id int) (*Deck, *Response, error) {
	// TODO: Use +id+
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("/api/v1/decks/%d", id),
		nil,
	)

	if err != nil {
		return nil, nil, err
	}

	f := &Deck{}
	resp, err := s.client.Do(req, f)

	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
