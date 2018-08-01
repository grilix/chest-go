package chest

type CardService struct {
	client *Client
}

type Card struct {
	Id    int    `json:"card_id"`
	Name  string `json:"card_name"`
	EditionCode string `json:"edition_code"`
	Count int    `json:"card_count"`
}
