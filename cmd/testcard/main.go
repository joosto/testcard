package main

import (
	"log"
	"github.com/joosto/stripe-test/pkg/model"
	"fmt"
	"flag"
	"os"
	"errors"
	"github.com/joosto/stripe-test/pkg/constants"
	"io/ioutil"
	"encoding/json"
)

func main() {
	l := log.New(os.Stderr, "", 0)

	displaySimple := flag.Bool("s", false, "simple")
	flag.Parse()
	queryCountry := flag.Arg(0)

	cards, err := retrieveCards(constants.CardFileName)
	if err != nil {
		l.Fatalf("failed to retrieve cards: %s", err)
	}
	foundCard := findCardForCountry(cards, queryCountry)
	displayCard(foundCard, *displaySimple)
}

func retrieveCards(fileName string) ([]*model.Card, error) {
	var cards []*model.Card

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not open %s", fileName))
	}
	err = json.Unmarshal(dat, &cards)
	if err != nil {
		return nil, errors.New("could not unmarshal")
	}

	return cards, nil
}

func displayCard(card *model.Card, displaySimple bool) {
	if displaySimple {
		fmt.Println(card.Number)
	} else {
		fmt.Println(card)
	}
}

func findCardForCountry(cards []*model.Card, queryCountry string) *model.Card {
	var found *model.Card

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
