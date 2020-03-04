package coinbasepro

import (
	"fmt"
)

type StablecoinConversion struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Amount        string `json:"amount"`
	// Response Fields
	ID            string `json:"id"`
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
}

func (c *Client) CreateConversion(newConversion *StablecoinConversion) (StablecoinConversion, error) {
	var savedConversion StablecoinConversion

	url := fmt.Sprintf("/conversions")
	_, err := c.Request("POST", url, newConversion, &savedConversion)

	return savedConversion, err
}
