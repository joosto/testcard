package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"flag"
)

func main() {
	displaySimple := flag.Bool("s", false, "simple")
	flag.Parse()
	queryCountry := flag.Arg(0)

	cards := RetrieveCards()
	foundCard := FindCardForCountry(cards, queryCountry)
	DisplayCard(foundCard, *displaySimple)
}
func RetrieveCards() []*Card {
	res, err := http.Get("https://stripe.com/docs/testing")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Page returned status %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cards []*Card

	s := doc.Find("tbody")
	s.Nodes = s.Nodes[2:5]
	s.Each(func(i int, tbody *goquery.Selection) {
		tbody.Find("tr").Each(func(j int, tr *goquery.Selection) {
			number := tr.Children().Eq(0).Text()
			token := tr.Children().Eq(1).Text()
			country := tr.Children().Eq(2).Text()

			countryParts := strings.Split(country, "(")
			countryName := countryParts[0]
			countryCode := countryParts[1][0:2]

			cards = append(cards, &Card{
				number,
				token,
				countryName,
				countryCode,
			})
		})
	})

	return cards
}

func DisplayCard(card *Card, displaySimple bool) {
	if displaySimple {
		fmt.Println(card.Number)
	} else {
		fmt.Println(card)
	}
}

func FindCardForCountry(cards []*Card, queryCountry string) *Card {
	var found *Card

	for _, card := range cards {
		if card.CountryCode == queryCountry {
			found = card
		}
	}

	if found == nil {
		log.Fatalf("No test card for country %s", queryCountry)
	}

	return found
}

type Card struct {
	Number      string
	Token       string
	Country     string
	CountryCode string
}

func (c *Card) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", c.Number, c.Token, c.Country, c.CountryCode)
}
