package selector

import (
	"errors"
	"fmt"
	"log"
	"github.com/Poleron402/Polinka/database"
	"github.com/charmbracelet/huh"
)
func listDecks() ([]polinkadatabase.Deck, error) {
	decks, err := polinkadatabase.ListDecks()
	if err != nil {
		fmt.Printf("%v", err)
		return decks, err
	}
	return decks, nil
}

func manageError(specialMessage string, errorString string) {
	fmt.Printf("%s\n=========================\n%v", specialMessage, errorString)
}
func SelectAction() {
	const (
		deck_create string = "deck_create"
		flashcard_practice string = "flashcard_practice"
		delete_deck string = "delete_deck"
		deck_practice string = "deck_practice" 
		deck_list string = "deck_list"
	)
	for {
		
		var option string
		form := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().Title("Where would you like to start? (Hold 'Ctrl+C' to quit)").Options(
				huh.NewOption("Create a new Deck", deck_create),
				huh.NewOption("Add a new flashcard to a Deck", flashcard_practice),
				huh.NewOption("Delete a Deck", delete_deck),
				huh.NewOption("Practice Deck", deck_practice),
				huh.NewOption("See All Decks", deck_list),
			).Value(&option),
		))
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		if option == deck_create {
			var deckName string
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name the deck").Value(&deckName).Validate(func(str string) error {
						decks, err := listDecks()
						if err != nil {
							fmt.Printf("%v", err)
						}
						for _, deck := range decks {
							if deck.Name == deckName{
								new_error := fmt.Sprintf("A deck with this name already exists! Pick a different name; %v", err)
								manageError("", new_error)
								return errors.New(new_error)
							}
						}
						return nil
					}),
				))
			err := form.Run()
			if err != nil {
				manageError("There has been an error running the form", err.Error())
			}
			err = polinkadatabase.CreateDeck(deckName)
			if err != nil {
				manageError("There has been an error creating a deck with this name", err.Error())
			}
		} else if option == deck_list {
			decks, err := listDecks()
			if err != nil{
				manageError("There has been an error fetching the Decks", err.Error())
			}
			fmt.Println("\nCurrent decks:\n===============================")
			for _, deck := range decks {
				fmt.Println(deck.Name)
			}
			fmt.Println("\n===============================")

		} else if option == delete_deck {
			deleteDeck()
		}
	}
}

func deleteDeck() {
	decks, err := listDecks()
	if err != nil {
		manageError("There was an issue listing decks", err.Error())
		return
	}
	options := make([]huh.Option[string], 0, len(decks))

	for _, deck := range decks {
		options = append(options, huh.NewOption(deck.Name, deck.Name))
	}
	var deckForDeletion []string
	form := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[string]().Title("Select Deck(s) to DELETE").Options(
			options...
		).Value(&deckForDeletion),
	))
	err = form.Run()
	if err!=nil{
		manageError("There has been an issue with the Deck selection form", err.Error())
	}
	for _, deck := range deckForDeletion {
		err := polinkadatabase.DeleteDeck(deck)
		if err != nil {
			manageError("There has been an issue with the Deck  deletion", err.Error())
		}
	}
}

func selectDeck() {
	
}