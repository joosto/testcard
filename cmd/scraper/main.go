package main

import (
	"github.com/joosto/stripe-test/pkg/model"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"errors"
	"fmt"
	"os"
	"encoding/json"
	"bytes"
	"github.com/joosto/stripe-test/pkg/constants"
)

func main() {
	l := log.New(os.Stderr, "", 0)

	cards, err := scrapeCards(constants.StripeDocsUrl)
	if err != nil {
		l.Fatalf("failed to retrieve cards: %s", err)
	}

	err = saveCards(cards, constants.CardFileName)
	if err != nil {
		l.Fatalf("failed to save cards: %s", err)
	}
}
func saveCards(cards []*model.Card, cacheFileName string) error {
	f, err := os.OpenFile(cacheFileName, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("could not open %s", cacheFileName))
	}

	b, err := json.Marshal(cards)
	if err != nil {
		errors.New("could not marshal the cards")
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")

	_, err = out.WriteTo(f)
	if err != nil {
		errors.New(fmt.Sprintf("could not write to %s", constants.CardFileName))
	}

	return nil
}

func scrapeCards(url string) ([]*model.Card, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not reach %s, please check your connection", url))
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s responded with %d %s", url, res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cards []*model.Card

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

			cards = append(cards, &model.Card{
				Number:      number,
				Token:       token,
				Country:     countryName,
				CountryCode: countryCode,
			})
		})
	})

	return cards, nil
}
