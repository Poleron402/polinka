package polinkadatabase

import (
	"database/sql"
	"log"
"fmt"
	_ "github.com/mattn/go-sqlite3"
)
var db *sql.DB

type Deck struct {
	ID int64
	Name string
}
type Flashcard struct {
	Question string
	Answer string
	Hint string
}
func CreateDBIfNotExist() {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()

	sqlStatement := `create table if not exists polinka_deck (id integer not null primary key autoincrement,
	name text unique);`

	sqlStatement1 := `create table if not exists polinka_flashcard (id integer not null primary key autoincrement,
	question text,
	answer text,
	hint text,
	deck_id integer,
	foreign key(deck_id) references polinka_deck(id)
	);`

	// Execute and check the first statement
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	// Execute and check the second statement
	_, err = db.Exec(sqlStatement1)
	if err != nil {
		log.Fatal(err)
	}
}
func CreateDeck(name string) error {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec(`insert into polinka_deck (name) values (?);`, name) 
	if err != nil {
		return err
	}
	return nil
}

func ListDecks() ([]Deck, error) {
	var decks []Deck
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	rows, err := db.Query(`select id, name from polinka_deck`)
	if err != nil {
		log.Fatal("There was an error fetching decks.")
		return decks, err
	}
	for rows.Next() {
		var deck Deck
		if err:=rows.Scan(&deck.ID, &deck.Name); err != nil {
			log.Fatal("There was an error parsing decks to an object.")
		}
		decks = append(decks, deck)
	}
	return decks, nil
}
func DeleteDeck(name string) error {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	_, err = db.Exec(`delete from polinka_deck where name=(?)`, name)
	if err != nil {
		return err
	}
	return nil
}
func GetFlashCards(deckId string) ([]Flashcard, error){
	var flashcards []Flashcard
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	rows, err := db.Query(`select question, answer, hint from polinka_flashcard where deck_id=(?)`, deckId)
	if err != nil {
		return flashcards, err
	}

	for rows.Next(){
		var fCard Flashcard
		if err:=rows.Scan(&fCard.Question, &fCard.Answer, &fCard.Hint); err != nil {
			log.Fatal("There was an error parsing flashcards to an object.")
		}
		flashcards = append(flashcards, fCard)
	}

	return flashcards, nil
}

func AddFlashcardsToDeck(question string, answer string, hint string, deckId string) error {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	_, err = db.Exec(`insert into polinka_flashcard (question, answer, hint, deck_id) values (?, ?, ?, ?);`, question, answer, hint, deckId) 
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	return nil
}