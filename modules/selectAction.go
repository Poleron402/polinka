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
		flashcard_add string = "flashcard_add"
		delete_deck string = "delete_deck"
		deck_practice string = "deck_practice" 
		deck_list string = "deck_list"
	)
	for {
		
		var option string
		form := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().Title("Where would you like to start? (Hold 'Ctrl+C' to quit)").Options(
				huh.NewOption("Create a new Deck", deck_create),
				huh.NewOption("Add a new flashcard to a Deck", flashcard_add),
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
		} else if option == deck_practice {
			selectDeck()
		}else if option == flashcard_add {
			addFlashcardsToDeck()
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

func listDecksForm(message string, selectionType string) (string, error) {
	decks, err := listDecks()
	if err != nil {
		manageError("There was an issue listing decks", err.Error())
		return "", err
	}
	if len(decks) == 0 {
		fmt.Println("There are no Decks yet. Create a new deck in the main menu.")
		return "", err
	}
	options := make([]huh.Option[string], 0, len(decks))

	for _, deck := range decks {
		if selectionType == "DeckName" {
			options = append(options, huh.NewOption(deck.Name, deck.Name))
		}else {
			options = append(options, huh.NewOption(deck.Name, fmt.Sprintf("%v",deck.ID)))
		}
	}
	var practiceDeck string
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().Title(message).Options(
			options...
		).Value(&practiceDeck),
	)).WithTheme(huh.ThemeCatppuccin())
	err = form.Run()
	if err!=nil{
		manageError("There has been an issue with the Deck selection form", err.Error())
		return "", err
	}
	return practiceDeck, nil
}

func selectDeck() {
	practiceDeck, err := listDecksForm("Select a Deck to practice", "DeckID")

	if err != nil {
		return
	}
	var flashcards []polinkadatabase.Flashcard
	flashcards, err = polinkadatabase.GetFlashCards(practiceDeck)
	if err != nil {
		manageError("There has been an issue fetching Flashcards from Deck", err.Error())
		return 
	}
	if len(flashcards) == 0 {
		manageError("There are no flashcards in this deck yet. Populate them in the main menu", "")
		return
	}


}

func addFlashcardsToDeck() {
	deckToAdd, err := listDecksForm("Select a Deck to practice", "DeckID")
	if err != nil {
		manageError("There has been an issue fetching decks.", err.Error())
		return
	}
	var (
		question string 
		answer string 
		hint string
	)
	huh.NewInput().
    Title("Front side (Question)").
    Value(&question).
    Run()
	huh.NewInput().
    Title("Back side (Answer)").
    Value(&answer).
    Run()
	huh.NewInput().
    Title("Hint.").
    Value(&hint).
    Run()

	err = polinkadatabase.AddFlashcardsToDeck(question, answer, hint, deckToAdd)
	if err != nil {
		manageError("There has been an issue adding cards to the deck.", err.Error())
		return
	}
	
}