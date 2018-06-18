package model

import "fmt"

type Card struct {
	Number      string `json:"number"`
	Token       string `json:"token"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

func (c *Card) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", c.Number, c.Token, c.Country, c.CountryCode)
}
