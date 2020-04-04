package coinbasepro

import (
	"fmt"
)

type Currency struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	MinSize       string   `json:"min_size"`
	// Used in WebSocket Feed "status" channel
	Status        string   `json:"status"`
	StatusMessage string   `json:"status_message"`
	MaxPrecision  string   `json:"max_precision"`
	ConvertibleTo []string `json:"convertible_to"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}
