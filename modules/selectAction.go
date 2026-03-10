package selector

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
	"unicode/utf8"

	polinkadatabase "github.com/Poleron402/Polinka/database"
	"github.com/charmbracelet/huh"
	"github.com/eiannone/keyboard"
	"github.com/mattn/go-runewidth"
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
		deck_create   string = "deck_create"
		flashcard_add string = "flashcard_add"
		delete_deck   string = "delete_deck"
		deck_practice string = "deck_practice"
		deck_list     string = "deck_list"
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
							if deck.Name == deckName {
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
			if err != nil {
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
			practiceDeckCards()
		} else if option == flashcard_add {
			fmt.Println("Click Ctrl+C to cancel")
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
		huh.NewMultiSelect[string]().Title("Select Deck(s) to DELETE (Press Space bar to select)").Options(
			options...,
		).Value(&deckForDeletion),
	))
	err = form.Run()
	if err != nil {
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
		} else {
			options = append(options, huh.NewOption(deck.Name, fmt.Sprintf("%v", deck.ID)))
		}
	}
	var practiceDeck string
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().Title(message).Options(
			options...,
		).Value(&practiceDeck),
	)).WithTheme(huh.ThemeCatppuccin())
	err = form.Run()
	if err != nil {
		manageError("There has been an issue with the Deck selection form", err.Error())
		return "", err
	}
	return practiceDeck, nil
}

func practiceDeckCards() {
	practiceDeck, err := listDecksForm("Select a Deck to practice", "DeckID")
	if err != nil {
		manageError("There has been an issue selecting the Deck", err.Error())
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
	err = keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// shuffle slice
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(flashcards), func(i, j int) {
		flashcards[i], flashcards[j] = flashcards[j], flashcards[i]
	})
	for i := 0; i < len(flashcards); i++ {
		card := flashcards[i]
		stopPractice := false
		questionShown := false
		if len(card.Question) > 0 {
			for {
				if !questionShown {
					questionShown = true
					splitcard := splitByRunecount(card.Question)
					fmt.Print("\033[H\033[2J")
					constructedCard := showFlashCard(splitcard)

					fmt.Print(constructedCard)
				}
				char, key, err := keyboard.GetKey()
				if err != nil {
					panic(err)
				}

				if len(card.Question) == 0 {
					continue
				}
				shallBreak := false
				switch key {
					case keyboard.KeyArrowRight:
						splitcard := splitByRunecount(card.Answer)
						constructedCard := showFlashCard(splitcard)
						fmt.Print("\033[H\033[2J")
						fmt.Println(constructedCard)

					case keyboard.KeyArrowLeft:
						splitcard := splitByRunecount(card.Question)
						constructedCard := showFlashCard(splitcard)
						fmt.Print("\033[H\033[2J")
						fmt.Println(constructedCard)
					case keyboard.KeyArrowUp:
						if i > 0 {
							i--
							card = flashcards[i]
							splitcard := splitByRunecount(card.Question)
							constructedCard := showFlashCard(splitcard)
							fmt.Print("\033[H\033[2J")
							fmt.Println(constructedCard)
						} else {
							splitcard := splitByRunecount(card.Question)
							constructedCard := showFlashCard(splitcard)
							fmt.Print("\033[H\033[2J")
							fmt.Println(constructedCard)
						}
					case keyboard.KeyEnter:
						shallBreak = true
						break
					case keyboard.KeyArrowDown:
						shallBreak = true
						break
					case keyboard.KeyEsc:
						stopPractice = true
						break
					default:
						break
				}
				if shallBreak || stopPractice {
					break
				}
				if char == 'h' {
					splitcard := splitByRunecount(card.Hint)
					constructedCard := showFlashCard(splitcard)
					fmt.Print("\033[H\033[2J")
					fmt.Println(constructedCard)
				}
			}
		}
		if stopPractice {
			fmt.Print("\n\n... Ending Practice ...\n\n")
			break
		}
	}
}
// func printFlashCard(typeOfCard string) {
// 	fmt.Print("\033[H\033[2J")
// 	splitcard := splitByRunecount(typeOfCard)
// 	constructedCard := showFlashCard(splitcard)
// 	fmt.Println(constructedCard)
// }

// https://thevalleyofcode.com/lesson/go-basics/go-tutorial-cowsay/ will copy to build a card
func showFlashCard(s []string) string {

	border := "░"
	max_length := 60
	bordered_length := 58
	padding := 2
	card := "\n\n"
	for i := 0; i < max_length; i++ {
		card += border
	}
	card += "\n"
	for range padding {
		card += "░"
		for range bordered_length {
			card += " "
		}
		card += "░\n"
	}

	for i := 0; i < len(s); i++ {
		card += "░ "
		card += s[i]
		card += " ░\n"
	}

	for range padding {
		card += "░"
		for range bordered_length {
			card += " "
		}
		card += "░\n"
	}
	for i := 0; i < max_length; i++ {
		card += border
	}
	card += "\n\n\nPress\n\x1b[34m↓ or Enter\x1b[0m for next card,\n\x1b[32m→\x1b[0m for Answer,\n\x1b[32m←\x1b[0m for back to Question,\n\x1b[33m'h'\x1b[0m for Hint,\n\x1b[35m'↑'\x1b[0m to go back to previous card,\n\x1b[94m'Esc'\x1b[0m to quit\n\n"
	return card
}

func splitByRunecount(s string) []string {
	// For Japanese language (and, perhaps, future foreign languages as well, RuneCountInString needs to be implemented in this function)
	var chunks []string
	runes := 56
	count := runewidth.StringWidth(s)
	if count < runes {
		for count < runes {
			s = " " + s
			count++
			if count == runes {
				break
			}
			s += " "
			count++
		}
		chunks = append(chunks, s)
	} else {
		for utf8.RuneCountInString(s) > 0 {
			j := 0
			for i := 0; i < runes && j < len(s); i++ {
				_, size := utf8.DecodeRuneInString(s[i:])
				j += size
			}
			chunks = append(chunks, s[:j])
			s = s[j:]
		}
		lastLine := &chunks[len(chunks)-1]
		if len(*lastLine) < runes {
			for len(*lastLine) < runes {
				*lastLine = " " + *lastLine
				if len(*lastLine) == runes {
					break
				}
				*lastLine += " "
			}
		}
	}
	return chunks
}
func addFlashcardsToDeck() {
	deckToAdd, err := listDecksForm("Select a Deck to populate with flashcards", "DeckID")
	if err != nil {
		manageError("There has been an issue fetching decks.", err.Error())
		return
	}
	var (
		question string
		answer   string
		hint     string
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

	if len(question) > 0 && len(answer) > 0 {
		err = polinkadatabase.AddFlashcardsToDeck(question, answer, hint, deckToAdd)
		if err != nil {
			manageError("There has been an issue adding cards to the deck.", err.Error())
			return
		}
	}
}
