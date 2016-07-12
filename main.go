package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MOZGIII/anki-read/anki"
)

func main() {
	defer anki.CleanupSQLiteTmpFiles()

	var cards []anki.Card

	for _, pkg := range os.Args[1:] {
		log.Printf("Reading from %s", pkg)

		p, err := anki.ReadFile(pkg)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		c, err := p.Collection.DB.Cards()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		cards = append(cards, c...)
	}

	for _, card := range cards {
		fmt.Printf("%s;%s;%s\n", card.Jap1, card.Jap2, card.Desc)
	}
}
